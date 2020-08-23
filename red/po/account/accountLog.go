package account

import (
	"github.com/segmentio/ksuid"
	"gitlab.kay.com/red/common"
	"gitlab.kay.com/red/proto/account"
	"time"
)

type AccountLog struct {
	Id              int64  `db:"id,omitempty"`         //
	LogNo           string `db:"log_no,uni"`           //流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string `db:"trade_no"`             //交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string `db:"account_no"`           //账户编号 账户ID
	AccountType     int    `db:"account_no"`           //账户编号 账户ID
	UserId          uint64 `db:"user_id"`              //用户编号
	UserName        string `db:"username"`             //用户名称
	TargetAccountNo string `db:"target_account_no"`    //账户编号 账户ID
	TargetUserId    uint64 `db:"target_user_id"`       //目标用户编号
	TargetUserName  string `db:"target_username"`      //目标用户名称
	Amount          uint64 `db:"amount"`               //交易金额,该交易涉及的金额
	ChangeType      int8   `db:"change_type"`          //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag      int8   `db:"change_flag"`          //交易变化标识:1:进账 2:出账 枚举
	Status          int8   `db:"status"`               //交易状态：
	Decs            string `db:"decs"`                 //交易描述
	CreatedAt       uint64 `db:"created_at,omitempty"` //创建时间
}

func (po *AccountLog) FromTransferPROTO(proto *red_proto_account.AccountTransfer) {
	po.createAccountLogNo()
	po.createTradeNo()

	po.AccountNo = proto.TradeBody.AccountNo
	po.AccountType = int(proto.TradeBody.AccountType)
	po.UserId = proto.TradeBody.UserId
	po.UserName = proto.TradeBody.UserName

	po.TargetAccountNo = proto.TradeTarget.AccountNo
	po.TargetUserId = proto.TradeTarget.UserId
	po.TargetUserName = proto.TradeTarget.UserName

	po.Amount = proto.Amount
	po.ChangeType = int8(proto.ChangeType)
	po.ChangeFlag = int8(proto.ChangeFlag)
	po.Status = 1
	po.Decs = proto.Desc
	po.CreatedAt = uint64(time.Now().Unix())
}

func (po *AccountLog) FromAccount(account *Account) {
	//通过account来创建流水，创建账户逻辑在前
	po.createAccountLogNo()
	po.createTradeNo()
	po.TradeNo = po.LogNo
	po.AccountType = account.AccountType
	// 交易主体
	po.AccountNo = account.AccountNo
	po.UserId = account.UserId
	po.UserName = account.UserName

	// 交易对象
	po.Amount = account.Balance
	po.TargetAccountNo = account.AccountNo
	po.TargetUserId = account.UserId
	po.TargetUserName = account.UserName

	// 交易类型
	po.ChangeType = common.FlagTransferIn
	po.ChangeFlag = 1

	po.Status = 1
	po.Decs = "创建账户"
	po.CreatedAt = uint64(time.Now().Unix())
}

func (po *AccountLog) createAccountLogNo() {
	po.LogNo = ksuid.New().Next().String()
}

func (po *AccountLog) createTradeNo() {
	po.TradeNo = ksuid.New().Next().String()
}
