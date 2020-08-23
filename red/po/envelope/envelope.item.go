package envelope

import (
	"github.com/segmentio/ksuid"
	"gitlab.kay.com/red/proto/envelope"
)

type RedEnvelopeItem struct {
	Id           uint64 `xorm:"autoincr"`
	ItemNo       string
	EnvelopeNo   string
	RecvUserName string
	RecvUserId   uint64
	RecvAccountNo    string
	Amount       uint64
	Quantity     uint64
	PayStatus    uint8
	CreatedAt    uint64
	UpdatedAt    uint64
	Desc         string
}

func (po *RedEnvelopeItem) ToPROTO() *red_proto_envelope.RedEnvelopeItem {
	envelopeItem := &red_proto_envelope.RedEnvelopeItem{
		ItemNo:       po.ItemNo,
		EnvelopeNo:   po.EnvelopeNo,
		UserName:     po.RecvUserName,
		UserId:       po.RecvUserId,
		Amount:       po.Amount,
		Quantity:     po.Quantity,
		AccountNo:    po.RecvAccountNo,
		PayStatus:    uint32(po.PayStatus),
		CreatedAt:    po.CreatedAt,
		UpdateAt:     po.UpdatedAt,
		Desc:         po.Desc,
	}
	return envelopeItem
}

func (po *RedEnvelopeItem) FromPROTO(envelopeItem *red_proto_envelope.RedEnvelopeItem) {
	po.ItemNo = envelopeItem.ItemNo
	po.EnvelopeNo = envelopeItem.EnvelopeNo
	po.RecvUserName = envelopeItem.UserName
	po.RecvUserId = envelopeItem.UserId
	po.Amount = envelopeItem.Amount
	po.Quantity = envelopeItem.Quantity
	po.RecvAccountNo = envelopeItem.AccountNo
	po.PayStatus = uint8(envelopeItem.PayStatus)
	po.CreatedAt = envelopeItem.CreatedAt
	po.UpdatedAt = envelopeItem.UpdateAt
	po.Desc = envelopeItem.Desc
}

func (po *RedEnvelopeItem) CreateEnvelopeItemNo()  {
	po.ItemNo = ksuid.New().Next().String()
}
