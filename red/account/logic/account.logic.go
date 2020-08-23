package logic

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/common"
	"gitlab.kay.com/red/po/account"
	"gitlab.kay.com/red/proto/account"
	"time"
	"xorm.io/core"
)

type AccountLogic struct {
	engine *xorm.Engine
}

func NewAccountLogic() (*AccountLogic, error)  {
	logic := new(AccountLogic)

	engine, err := xorm.NewEngine("mysql", config.Get("mysql", "xproject"))
	if err != nil {
		logger.Errorf("xorm NewEngine err:%s", err.Error())
		return nil, err
	}

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	fieldMapper := core.NewPrefixMapper(core.SnakeMapper{}, "f")

	engine.SetTableMapper(tbMapper)
	engine.SetColumnMapper(fieldMapper)

	logic.engine = engine
	return logic, nil
}

func (logic *AccountLogic) CreateAccount (ctx context.Context, req *red_proto_account.CreateAccountRequest, resp *red_proto_account.CreateAccountResponse) error {
	logger.Infof("logic CreateAccount start req:%s", req.String())

	header := new(red_proto_account.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RespHeader = header

	sess := logic.engine.NewSession()
	ha := new(account.Account)
	ok, err := sess.Table("t_account").Where("fuser_id=? and faccount_type=?",
		req.GetUserId(), req.GetAccountType()).Get(ha)
	if err != nil {
		logger.Errorf("CreateAccount Insert err:%s", err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	if ok {
		logger.Errorf("acount has exist account:%v", ha)
		header.Code = common.ERR_HAS_ACCOUNT.GetErrCode()
		header.Msg = common.ERR_HAS_ACCOUNT.GetErrMsg()
		return nil
	}

	a := new(account.Account)
	a.UserId = req.GetUserId()
	a.UserName = req.GetUserName()
	a.AccountName = req.GetAccountName()
	a.AccountType = int(req.GetAccountType())
	a.CurrencyCode = req.GetCurrencyCode()
	a.Balance = req.GetAmount()
	a.CreateAccountNo()
	a.SetIsValid()
	a.CreatedAt = uint64(time.Now().Unix())
	a.UpdatedAt = uint64(time.Now().Unix())

	accountLog := new(account.AccountLog)
	accountLog.FromAccount(a)

	_, err = sess.Table("t_account").Insert(a)
	if err != nil {
		logger.Errorf("CreateAccount Insert err:%s", err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	_, err = sess.Table("t_account_log").Insert(accountLog)
	if err != nil {
		logger.Errorf("CreateAccount Insert err:%s", err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	resp.Account = a.ToPROTO()
	logger.Infof("logic CreateAccount success")
	return nil
}

func (logic *AccountLogic) Transfer (ctx context.Context, req *red_proto_account.AccountTransferRequest, resp *red_proto_account.AccountTransferResponse) error {
	logger.Infof("logic Transfer start req:%s", req.String())

	header := new(red_proto_account.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RespHeader = header

	transfer := req.GetAccountTransfer()

	if transfer.Amount == 0 {
		resp.TransferStatus = common.TransferFailed
		logger.Errorf("Transfer param failed transfer:%v", transfer)
		header.Code = common.ERR_PARAM.GetErrCode()
		header.Msg = common.ERR_PARAM.GetErrMsg()
		return nil
	}

	if transfer.ChangeType == common.FlagTransferOut {
		if transfer.ChangeFlag > 0 {
			resp.TransferStatus = common.TransferFailed
			logger.Errorf("input param err req:%v", req)
			header.Code = common.ERR_PARAM.GetErrCode()
			header.Msg = common.ERR_PARAM.GetErrMsg()
			return nil
		}
	}else {
		if transfer.ChangeFlag < 0 {
			resp.TransferStatus = common.TransferFailed
			logger.Errorf("input param err req:%v", req)
			header.Code = common.ERR_PARAM.GetErrCode()
			header.Msg = common.ERR_PARAM.GetErrMsg()
			return nil
		}
	}

	errMsg := logic.realTransfer(transfer)
	if errMsg.GetErrCode() != 0 {
		resp.TransferStatus = common.TransferFailed
		logger.Errorf("Transfer realTransfer failed transfer:%v", transfer)
		header.Code = errMsg.GetErrCode()
		header.Msg = errMsg.GetErrMsg()
		return nil
	}

	tradeBody := transfer.TradeBody
	tradeTarget := transfer.TradeTarget

	otherTransfer := transfer
	otherTransfer.TradeBody = tradeTarget
	otherTransfer.TradeTarget = tradeBody
	otherTransfer.ChangeFlag = -transfer.ChangeFlag
	otherTransfer.ChangeType = -transfer.ChangeType

	errMsg = logic.realTransfer(otherTransfer)
	if errMsg.GetErrCode() != 0 {
		resp.TransferStatus = common.TransferFailed
		logger.Errorf("Transfer realTransfer failed transfer:%v", transfer)
		header.Code = errMsg.GetErrCode()
		header.Msg = errMsg.GetErrMsg()
		return nil
	}

	resp.TransferStatus = common.TransferSuccess
	logger.Infof("logic Transfer success")
	return nil
}

func (logic *AccountLogic) StoreValue (ctx context.Context, req *red_proto_account.StoreValueRequest, resp *red_proto_account.StoreValueResponse) error {
	logger.Infof("logic StoreValue start req:%s", req.String())

	header := new(red_proto_account.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RespHeader = header

	transfer := req.GetAccountTransfer()
	if transfer.Amount == 0 {
		resp.TransferStatus = common.TransferFailed
		logger.Errorf("StoreValue param failed transfer:%v", transfer)
		header.Code = common.ERR_PARAM.GetErrCode()
		header.Msg = common.ERR_PARAM.GetErrMsg()
		return nil
	}

	transfer.TradeTarget = transfer.TradeBody
	errMsg := logic.realTransfer(transfer)
	if errMsg.GetErrCode() != 0 {
		resp.TransferStatus = common.TransferFailed
		logger.Errorf("StoreValue realTransfer failed transfer:%v", transfer)
		header.Code = errMsg.GetErrCode()
		header.Msg = errMsg.GetErrMsg()
		return nil
	}

	resp.TransferStatus = common.TransferSuccess
	logger.Infof("logic StoreValue success")
	return nil
}

func (logic *AccountLogic) GetEnvelopeAccountByUserId (ctx context.Context, req *red_proto_account.GetEnvelopeAccountByUserIdRequest, resp *red_proto_account.GetEnvelopeAccountByUserIdResponse) error {
	logger.Infof("logic GetEnvelopeAccountByUserId start req:%s", req.String())

	header := new(red_proto_account.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RespHeader = header

	userId := req.GetUserId()
	var ma []*account.Account

	sess := logic.engine.NewSession()
	err := sess.Table("t_account").Where("fuser_id=?", userId).Find(&ma)
	if err != nil {
		logger.Errorf("CreateAccount Insert err:%s", err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	var protoAccounts []*red_proto_account.Account
	for _, a := range ma {
		protoAccounts = append(protoAccounts, a.ToPROTO())
	}

	resp.Accounts = protoAccounts
	logger.Infof("logic GetEnvelopeAccountByUserId success")
	return nil
}

func (logic *AccountLogic) GetAccount (ctx context.Context, req *red_proto_account.GetAccountRequest, resp *red_proto_account.GetAccountResponse) error {
	logger.Infof("logic GetAccount start req:%s", req.String())

	header := new(red_proto_account.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RespHeader = header

	accountNo := req.GetAccountNo()
	a := new(account.Account)
	a.AccountNo = accountNo

	sess := logic.engine.NewSession()
	ok, err := sess.Table("t_account").Get(a)
	if err != nil {
		logger.Errorf("CreateAccount Insert err:%s", err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	if !ok {
		header.Code = common.ERR_HAS_NOT_ACCOUNT.GetErrCode()
		header.Msg = common.ERR_HAS_NOT_ACCOUNT.GetErrMsg()
		return nil
	}

	resp.Account = a.ToPROTO()
	logger.Infof("logic GetAccount success")
	return nil
}

func (logic *AccountLogic) realTransfer(transfer *red_proto_account.AccountTransfer) common.ErrCodeMsg {
	sess := logic.engine.NewSession()
	a := new(account.Account)
	a.AccountNo = transfer.TradeBody.AccountNo

	ok, err := sess.Table("t_account").Get(a)
	if err != nil {
		logger.Errorf("CreateAccount Insert err:%s", err.Error())
		return common.ERR_OPERAROR
	}

	if !ok {
		return common.ERR_HAS_NOT_ACCOUNT
	}

	if transfer.ChangeType == common.FlagTransferOut {
		if transfer.Amount > a.Balance {
			logger.Errorf("realTransfer Amount:%d Account:%v", transfer.Amount, a)
			return common.ERR_NOT_BALANCE
		}

		a.Balance = a.Balance - transfer.Amount

	}else {
		a.Balance = a.Balance + transfer.Amount
	}
	a.UpdatedAt = uint64(time.Now().Unix())
	_, err = sess.Table("t_account").Where("fid=?", a.Id).Update(a)
	if err != nil {
		logger.Errorf("CreateAccount Insert err:%s", err.Error())
		return common.ERR_OPERAROR
	}

	accountLog := new(account.AccountLog)
	accountLog.FromTransferPROTO(transfer)

	_, err = sess.Table("t_account_log").Insert(accountLog)
	if err != nil {
		logger.Errorf("CreateAccount Insert err:%s", err.Error())
		return common.ERR_OPERAROR
	}

	return common.ERR_SUCCESS
}

