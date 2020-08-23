package envelope

import (
	"github.com/segmentio/ksuid"
	"gitlab.kay.com/red/common"
	"gitlab.kay.com/red/common/algo"
	"gitlab.kay.com/red/proto/envelope"
)

type RedEnvelopeGoods struct {
	Id               uint64 `db:"id,omitempty" xorm:"autoincr"`         //自增ID
	EnvelopeNo       string `db:"envelope_no,uni"`      //红包编号,红包唯一标识
	EnvelopeType     uint8  `db:"envelope_type"`        //红包类型：普通红包，碰运气红包
	UserName         string `db:"username"`             //用户名称
	UserId           uint64 `db:"user_id"`              //用户编号, 红包所属用户
	Blessing         string `db:"blessing"`             //祝福语
	Amount           uint64 `db:"amount"`               //红包总金额
	AmountOne        uint64 `db:"amount_one"`           //单个红包金额，碰运气红包无效
	Quantity         uint64 `db:"quantity"`             //红包总数量
	RemainAmount     uint64 `db:"remain_amount"`        //红包剩余金额额
	RemainQuantity   uint64 `db:"remain_quantity"`      //红包剩余数量
	ExpiredAt        uint64 `db:"expired_at"`           //过期时间
	Status           int8   `db:"status"`               //红包状态：0红包初始化，1启用，2失效
	OrderType        int8   `db:"order_type"`           //订单类型：发布单、退款单
	PayStatus        int8   `db:"pay_status"`           //支付状态：未支付，支付中，已支付，支付失败
	CreatedAt        uint64 `db:"created_at,omitempty"` //创建时间
	UpdatedAt        uint64 `db:"updated_at,omitempty"` //更新时间
	OriginEnvelopeNo string `db:"origin_envelope_no"`   //原关联订单号
}

func (po *RedEnvelopeGoods) ToPROTO() *red_proto_envelope.EnvelopeGoods {
	proto := &red_proto_envelope.EnvelopeGoods{
		EnvelopeNo:       po.EnvelopeNo,
		EnvelopeType:     uint32(po.EnvelopeType),
		UserName:         po.UserName,
		UserId:           po.UserId,
		Blessing:         po.Blessing,
		Amount:           po.Amount,
		AmountOne:        po.AmountOne,
		Quantity:         uint32(po.Quantity),
		RemainAmount:     po.RemainAmount,
		RemainQuantity:   uint32(po.RemainQuantity),
		ExpireAt:         po.ExpiredAt,
		Status:           uint32(po.Status),
		OrderType:        uint32(po.OrderType),
		PayStatus:        uint32(po.PayStatus),
		CreatedAt:        po.CreatedAt,
		UpdateAt:         po.UpdatedAt,
		OriginEnvelopeNo: po.OriginEnvelopeNo,
	}
	return proto
}

func (po *RedEnvelopeGoods) FromPROTO(proto *red_proto_envelope.EnvelopeGoods) {
	po.EnvelopeNo = proto.EnvelopeNo
	po.EnvelopeType = uint8(proto.EnvelopeType)
	po.UserName = proto.UserName
	po.UserId = proto.UserId
	po.Blessing = proto.Blessing
	po.Amount = proto.Amount
	po.AmountOne = proto.AmountOne
	po.Quantity = uint64(proto.Quantity)
	po.RemainAmount = proto.RemainAmount
	po.RemainQuantity = uint64(proto.RemainQuantity)
	po.ExpiredAt = proto.ExpireAt
	po.Status = int8(proto.Status)
	po.OrderType = int8(proto.OrderType)
	po.PayStatus = int8(proto.PayStatus)
	po.CreatedAt = proto.CreatedAt
	po.UpdatedAt = proto.UpdateAt
	po.OriginEnvelopeNo = proto.OriginEnvelopeNo
}

func (po *RedEnvelopeGoods) CreateEnvelopeNo()  {
	po.EnvelopeNo = ksuid.New().Next().String()
}

func (po *RedEnvelopeGoods) NextAmount() uint64 {
	if po.RemainQuantity == 1 {
		return po.RemainAmount
	}

	if po.EnvelopeType == common.GeneralEnvelopeType {
		return po.AmountOne
	}else {
		amount := algo.SimpleRand(int64(po.RemainQuantity), int64(po.RemainAmount))
		return uint64(amount)
	}
}
