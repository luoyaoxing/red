package logic

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/AoClient"
	"gitlab.kay.com/red/common"
	"gitlab.kay.com/red/po/envelope"
	"gitlab.kay.com/red/proto/envelope"
	"math"
	"strconv"
	"time"
	"xorm.io/core"
)

type RedEnvelopeLogic struct {
	userCli *AoClient.UserAoClient
	accountCli *AoClient.AccountAoClient
	engine *xorm.Engine
}

func NewRedEnvelopeLogic() (*RedEnvelopeLogic, error) {
	logic := new(RedEnvelopeLogic)
	logic.userCli = AoClient.NewUserAoClient()
	logic.accountCli = AoClient.NewAccountAoClient()

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

func (logic *RedEnvelopeLogic) SendOut (ctx context.Context, req *red_proto_envelope.RedEnvelopeSendRequest, resp *red_proto_envelope.RedEnvelopeSendResponse) error {
	logger.Infof("logic SendOut start req:%s", req.String())

	header := new(red_proto_envelope.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RspHeader = header

	now := time.Now()
	uid := req.GetUserId()
	blessing := req.GetBlessing()
	envelopeType := req.GetEnvelopeType()

	acc, err := logic.accountCli.GetEnvelopeAccountByUserId(uid)
	if err != nil || acc == nil {
		logger.Errorf("accountCli GetEnvelopeAccountByUserId uid:%d err:%s", uid, err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	goods := new(envelope.RedEnvelopeGoods)
	goods.CreateEnvelopeNo()
	goods.EnvelopeType = uint8(req.GetEnvelopeType())
	goods.UserId = req.GetUserId()
	goods.UserName = req.GetUserName()
	goods.Blessing = req.GetBlessing()
	if blessing == "" {
		goods.Blessing = common.DefaultBlessing
	}
	goods.Quantity = uint64(req.GetQuantity())
	if envelopeType == common.GeneralEnvelopeType {
		goods.AmountOne = req.GetAmount()
		goods.Amount = goods.AmountOne * uint64(goods.Quantity)
	}else {
		goods.AmountOne = 0
		goods.Amount = req.GetAmount()
	}
	goods.RemainAmount = goods.Amount
	goods.RemainQuantity = goods.Quantity
	goods.ExpiredAt = uint64(now.Add(5 * 86400).Unix())
	goods.Status = int8(common.OrderCreate)
	goods.OrderType = int8(common.OrderTypeSending)
	goods.PayStatus = int8(common.PayNothing)
	goods.CreatedAt = uint64(now.Unix())
	goods.UpdatedAt = uint64(now.Unix())

	sess := logic.engine.NewSession()
	_, insertErr := sess.Table("t_red_envelope_goods").Insert(goods)
	gid := goods.Id
	logger.Debugf("SendOut insert gid:%d", gid)
	if insertErr != nil {
		logger.Errorf("SendOut Insert err:%s", insertErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	status := logic.accountCli.RedSendOutTransfer(acc, goods.ToPROTO())
	if status < 0 {
		logger.Errorf("SendOut Transfer failed")
		header.Code = common.ERR_TRANSFER_FAILED.GetErrCode()
		header.Msg = common.ERR_TRANSFER_FAILED.GetErrMsg()
		return nil
	}

	updateGoods := new(envelope.RedEnvelopeGoods)
	updateGoods.Status = int8(common.OrderSending)
	updateGoods.PayStatus = int8(common.Payed)
	updateGoods.UpdatedAt = uint64(time.Now().Unix())
	_, updateErr := sess.Table("t_red_envelope_goods").Where("fid=?", gid).Update(updateGoods)
	if updateErr != nil {
		logger.Errorf("SendOut Update err:%s", updateErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	newGoods := new(envelope.RedEnvelopeGoods)
	ok, getErr := sess.Table("t_red_envelope_goods").Where("fid=?", gid).Get(newGoods)
	if getErr != nil {
		logger.Errorf("SendOut Get err:%s", getErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	if !ok {
		logger.Errorf("查不到红包")
		header.Code = common.ERR_HAS_NOT_RED_ENVELOPE.GetErrCode()
		header.Msg = common.ERR_HAS_NOT_RED_ENVELOPE.GetErrMsg()
		return nil
	}

	resp.EnvelopeGoods = newGoods.ToPROTO()
	resp.Link = "http://daniel.host.com/v1/link/envelope/" + strconv.Itoa(int(gid))
	logger.Infof("logic SendOut success")
	return nil
}

func (logic *RedEnvelopeLogic) Receive (ctx context.Context, req *red_proto_envelope.RedEnvelopeReceiveRequest, resp *red_proto_envelope.RedEnvelopeReceiveResponse) error {
	logger.Infof("logic Receive start req:%s", req.String())

	header := new(red_proto_envelope.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RspHeader = header

	userId := req.GetUserId()
	userName := req.GetUserName()
	accountNo := req.GetAccountNo()
	envelopeNo := req.GetEnvelopeNo()

	acc, err := logic.accountCli.GetEnvelopeAccountByUserId(userId)
	if err != nil || acc == nil {
		logger.Errorf("accountCli GetEnvelopeAccountByUserId userId:%d err:%s", userId, err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	sess := logic.engine.NewSession()
	goods := new(envelope.RedEnvelopeGoods)
	ok, getErr := sess.Table("t_red_envelope_goods").Where("fenvelope_no=?", envelopeNo).Get(goods)
	if getErr != nil {
		logger.Errorf("Receive Get err:%s", getErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	if !ok {
		logger.Errorf("查不到红包")
		header.Code = common.ERR_HAS_NOT_RED_ENVELOPE.GetErrCode()
		header.Msg = common.ERR_HAS_NOT_RED_ENVELOPE.GetErrMsg()
		return nil
	}

	if goods.Quantity <= 0 || goods.Amount <= 0 {
		logger.Errorf("Receive goods:%v", goods)
		header.Code = common.ERR_NO_RED_ENVELOPE.GetErrCode()
		header.Msg = common.ERR_NO_RED_ENVELOPE.GetErrMsg()
		return nil
	}

	item := new(envelope.RedEnvelopeItem)
	item.CreateEnvelopeItemNo()
	item.EnvelopeNo = envelopeNo
	item.RecvUserId = userId
	item.RecvUserName = userName
	item.RecvAccountNo = accountNo
	item.Quantity = 1
	item.Amount = goods.NextAmount()
	item.PayStatus = uint8(common.PayNothing)
	item.CreatedAt = uint64(time.Now().Unix())
	item.UpdatedAt = uint64(time.Now().Unix())
	item.Desc = userName + "抢了" + goods.UserName + "的红包"

	_, insertErr := sess.Table("t_red_envelope_item").Insert(item)
	itemId := item.Id
	if insertErr != nil {
		logger.Errorf("Receive Insert err:%s", insertErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	status := logic.accountCli.RedReceiveTransfer(acc, item.ToPROTO())
	if status < 0 {
		logger.Errorf("Receive RedReceiveTransfer failed")
		header.Code = common.ERR_TRANSFER_FAILED.GetErrCode()
		header.Msg = common.ERR_TRANSFER_FAILED.GetErrMsg()
		return nil
	}

	updateItem := new(envelope.RedEnvelopeItem)
	updateItem.PayStatus = uint8(common.Payed)
	updateItem.UpdatedAt = uint64(time.Now().Unix())
	_, updateErr := sess.Table("t_red_envelope_item").Where("fid=?", itemId).Update(updateItem)
	if updateErr != nil {
		logger.Errorf("Receive t_red_envelope_item Update err:%s", updateErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	updateGoods := new(envelope.RedEnvelopeGoods)
	updateGoods.RemainQuantity = goods.Quantity - item.Quantity
	updateGoods.RemainAmount = goods.Amount - item.Amount
	updateGoods.UpdatedAt = uint64(time.Now().Unix())
	_, updateErr = sess.Table("t_red_envelope_goods").Where("fid=?", goods.Id).Update(updateGoods)
	if updateErr != nil {
		logger.Errorf("Receive t_red_envelope_goods Update err:%s", updateErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	newItem := new(envelope.RedEnvelopeItem)
	ok, getErr = sess.Table("t_red_envelope_item").Where("fid=?", itemId).Get(newItem)
	if getErr != nil {
		logger.Errorf("Receive Get t_red_envelope_item err:%s", getErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	if !ok {
		logger.Errorf("查不到红包记录")
		header.Code = common.ERR_HAS_NOT_RED_ENVELOPE_ITEM.GetErrCode()
		header.Msg = common.ERR_HAS_NOT_RED_ENVELOPE_ITEM.GetErrMsg()
		return nil
	}

	resp.RedEnvelopeItem = newItem.ToPROTO()
	logger.Infof("logic Receive success")
	return nil
}

func (logic *RedEnvelopeLogic) Refund (ctx context.Context, req *red_proto_envelope.RedEnvelopeRefundRequest, resp *red_proto_envelope.RedEnvelopeRefundResponse) error {
	logger.Infof("logic Refund start req:%s", req.String())
	logger.Infof("logic Refund success")
	return nil
}

func (logic *RedEnvelopeLogic) Get (ctx context.Context, req *red_proto_envelope.GetRedEnvelopRequest, resp *red_proto_envelope.GetRedEnvelopeResponse) error {
	logger.Infof("logic Get start req:%s", req.String())

	header := new(red_proto_envelope.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RspHeader = header

	envelopeNo := req.GetEnvelopeNo()
	sess := logic.engine.NewSession()
	goods := new(envelope.RedEnvelopeGoods)
	ok, getErr := sess.Table("t_red_envelope_goods").Where("fenvelope_no=?", envelopeNo).Get(goods)
	if getErr != nil {
		logger.Errorf("Receive Get err:%s", getErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	if !ok {
		logger.Errorf("查不到红包")
		header.Code = common.ERR_HAS_NOT_RED_ENVELOPE.GetErrCode()
		header.Msg = common.ERR_HAS_NOT_RED_ENVELOPE.GetErrMsg()
		return nil
	}

	resp.EnvelopeGoods = goods.ToPROTO()
	logger.Infof("logic Get success")
	return nil
}

func (logic *RedEnvelopeLogic) ListSent (ctx context.Context, req *red_proto_envelope.ListSentRequest, resp *red_proto_envelope.ListSentResponse) error {
	logger.Infof("logic ListSent start req:%s", req.String())

	header := new(red_proto_envelope.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RspHeader = header

	uid := req.GetUserId()
	page := req.GetPage()
	pageSize := req.GetSize()

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	sess := logic.engine.NewSession()
	num, cErr := sess.Table("t_red_envelope_goods").Where("fuser_id=?", uid).Count()
	if cErr != nil {
		logger.Errorf("ListSent Get err:%s", cErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	var redGoods []*envelope.RedEnvelopeGoods
	findErr := sess.Table("t_red_envelope_goods").Where("fuser_id=?", uid).Desc("fcreated_at").Limit(int(pageSize), int(offset)).Find(&redGoods)
	if findErr != nil {
		logger.Errorf("ListSent Find err:%s", cErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	var envelopeGoods []*red_proto_envelope.EnvelopeGoods
	for _, redGood := range redGoods {
		envelopeGoods = append(envelopeGoods, redGood.ToPROTO())
	}

	totalPage := math.Ceil(float64(num) / float64(pageSize))
	resp.TotalPage = uint64(totalPage)
	resp.EnvelopeGoods = envelopeGoods
	logger.Infof("logic ListSent success")
	return nil
}

func (logic *RedEnvelopeLogic) ListReceived (ctx context.Context, req *red_proto_envelope.ListReceivedRequest, resp *red_proto_envelope.ListReceivedResponse) error {
	logger.Infof("logic ListReceived start req:%s", req.String())

	header := new(red_proto_envelope.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RspHeader = header

	uid := req.GetUserId()
	page := req.GetPage()
	pageSize := req.GetSize()

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	sess := logic.engine.NewSession()
	num, cErr := sess.Table("t_red_envelope_item").Where("frecv_user_id=?", uid).Count()
	if cErr != nil {
		logger.Errorf("ListReceived Get err:%s", cErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	var items []*envelope.RedEnvelopeItem
	findErr := sess.Table("t_red_envelope_item").Where("frecv_user_id=?", uid).Desc("fcreated_at").Limit(int(pageSize), int(offset)).Find(&items)
	if findErr != nil {
		logger.Errorf("ListReceived Find err:%s", cErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	var envelopeItems []*red_proto_envelope.RedEnvelopeItem
	for _, item := range items {
		envelopeItems = append(envelopeItems, item.ToPROTO())
	}

	totalPage := math.Ceil(float64(num) / float64(pageSize))
	resp.TotalPage = uint64(totalPage)
	resp.EnvelopeItems = envelopeItems
	logger.Infof("logic ListReceived success")
	return nil
}

func (logic *RedEnvelopeLogic) ListReceivable (ctx context.Context, req *red_proto_envelope.ListReceivableRequest, resp *red_proto_envelope.ListReceivableResponse) error {
	logger.Infof("logic ListReceivable start req:%s", req.String())

	header := new(red_proto_envelope.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RspHeader = header

	page := req.GetPage()
	pageSize := req.GetSize()

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	sess := logic.engine.NewSession()
	num, cErr := sess.Table("t_red_envelope_goods").Count()
	if cErr != nil {
		logger.Errorf("ListSent Get err:%s", cErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	var redGoods []*envelope.RedEnvelopeGoods
	findErr := sess.Table("t_red_envelope_goods").Desc("fcreated_at").Limit(int(pageSize), int(offset)).Find(&redGoods)
	if findErr != nil {
		logger.Errorf("ListSent Find err:%s", cErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	var envelopeGoods []*red_proto_envelope.EnvelopeGoods
	for _, redGood := range redGoods {
		envelopeGoods = append(envelopeGoods, redGood.ToPROTO())
	}

	totalPage := math.Ceil(float64(num) / float64(pageSize))
	resp.TotalPage = uint64(totalPage)
	resp.EnvelopeGoods = envelopeGoods
	logger.Infof("logic ListReceivable success")
	return nil
}

func (logic *RedEnvelopeLogic) ListItems (ctx context.Context, req *red_proto_envelope.ListItemsRequest, resp *red_proto_envelope.ListItemsResponse) error {
	logger.Infof("logic ListItems start req:%s", req.String())

	header := new(red_proto_envelope.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RspHeader = header

	envelopeNo := req.GetEnvelopeNo()
	sess := logic.engine.NewSession()
	var items []*envelope.RedEnvelopeItem
	findErr := sess.Table("t_red_envelope_item").Where("fenvelope_no=?", envelopeNo).Desc("fcreated_at").Find(&items)
	if findErr != nil {
		logger.Errorf("ListReceived Find err:%s", findErr.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	var envelopeItems []*red_proto_envelope.RedEnvelopeItem
	for _, item := range items {
		envelopeItems = append(envelopeItems, item.ToPROTO())
	}

	resp.EnvelopeItems = envelopeItems
	logger.Infof("logic ListItems success")
	return nil
}


