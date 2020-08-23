package common

const (
	DefaultBlessing   = "恭喜发财，鸿富猪到！"
	DefaultTimeFormat = "2006-01-02.15:04:05"
)

//订单类型：发布单、退款单
type OrderType int

const (
	OrderTypeSending OrderType = 1
	OrderTypeRefund  OrderType = 2
)

//支付状态：未支付，支付中，已支付，支付失败
//退款：未退款，退款中，已退款，退款失败
type PayStatus int

const (
	PayNothing PayStatus = 1
	Paying     PayStatus = 2
	Payed      PayStatus = 3
	PayFailure PayStatus = 4
	//
	RefundNothing PayStatus = 61
	Refunding     PayStatus = 62
	Refunded      PayStatus = 63
	RefundFailure PayStatus = 64
)

//红包订单状态：创建、发布、过期、失效、过期退款成功，过期退款失败
type OrderStatus int

const (
	OrderCreate                  OrderStatus = 1
	OrderSending                 OrderStatus = 2
	OrderExpired                 OrderStatus = 3
	OrderDisabled                OrderStatus = 4
	OrderExpiredRefundSuccessful OrderStatus = 5
	OrderExpiredRefundFalured    OrderStatus = 6
)

//红包类型：普通红包，碰运气红包
type EnvelopeType int

const (
	GeneralEnvelopeType = 1
	LuckyEnvelopeType   = 2
)

var EnvelopeTypes = map[EnvelopeType]string{
	GeneralEnvelopeType: "普通红包",
	LuckyEnvelopeType:   "碰运气红包",
}

const (
	FlagTransferIn = 1
	FlagTransferOut = -1
)

//转账的类型：0=创建账户 >=1进账 <=- 支出
//type ChangeType int8

const (
	//账户创建
	AccountCreated = 0
	//储值
	AccountStoreValue = 1
	//红包资金的支出
	EnvelopeOutgoing = -2
	//红包资金的收入
	EnvelopeIncoming = 2
	//红包过期退款
	EnvelopExpiredRefund = 3
)

const (
	TransferSuccess = 1
	TransferFailed = -1
)