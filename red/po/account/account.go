package account

import (
	"github.com/segmentio/ksuid"
	"gitlab.kay.com/red/proto/account"
)

// 账户持久化对象
type Account struct {
	Id           int64  `db:"id,omitempty"`         //账户ID
	AccountNo    string `db:"account_no,uni"`       //账户编号,账户唯一标识
	AccountName  string `db:"account_name"`         //账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱
	AccountType  int    `db:"account_type"`         //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode string `db:"currency_code"`        //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	UserId       uint64 `db:"user_id"`              //用户编号, 账户所属用户
	UserName     string `db:"username"`             //用户名称
	Balance      uint64 `db:"balance"`              //账户可用余额
	Status       int    `db:"status"`               //账户状态，账户状态：0账户初始化，1启用，2停用
	CreatedAt    uint64 `db:"created_at,omitempty"` //创建时间
	UpdatedAt    uint64 `db:"updated_at,omitempty"` //更新时间
}

func (a *Account) ToPROTO() *red_proto_account.Account  {
	account := new(red_proto_account.Account)
	account.UserId = a.UserId
	account.UserName = a.UserName
	account.AccountNo = a.AccountNo
	account.AccountName = a.AccountName
	account.AccountType = uint32(a.AccountType)
	account.CurrencyCode = a.CurrencyCode
	account.Status = uint32(a.Status)
	account.CreatedAt = a.CreatedAt
	account.UpdateAt = a.UpdatedAt
	return account
}

func (a *Account) FromPROTO(account *red_proto_account.Account)  {
	a.UserId = account.UserId
	a.UserName = account.UserName
	a.AccountNo = account.AccountNo
	a.AccountName = account.AccountName
	a.AccountType = int(account.AccountType)
	a.CurrencyCode = account.CurrencyCode
	a.Status = int(account.Status)
	a.CreatedAt = account.CreatedAt
	a.UpdatedAt = account.UpdateAt
}

func (a *Account) CreateAccountNo()  {
	a.AccountNo = ksuid.New().Next().String()
}

func (a *Account) SetIsValid()  {
	a.Status = 1
}
