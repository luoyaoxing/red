// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envelope.proto

package red_proto_envelope

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EnvelopeGoods struct {
	EnvelopeNo           string   `protobuf:"bytes,1,opt,name=envelope_no,json=envelopeNo,proto3" json:"envelope_no,omitempty"`
	EnvelopeType         uint32   `protobuf:"varint,2,opt,name=envelope_type,json=envelopeType,proto3" json:"envelope_type,omitempty"`
	AccountNo            string   `protobuf:"bytes,3,opt,name=account_no,json=accountNo,proto3" json:"account_no,omitempty"`
	OriginEnvelopeNo     string   `protobuf:"bytes,4,opt,name=origin_envelopeNo,json=originEnvelopeNo,proto3" json:"origin_envelopeNo,omitempty"`
	UserName             string   `protobuf:"bytes,5,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserId               uint64   `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Blessing             string   `protobuf:"bytes,7,opt,name=blessing,proto3" json:"blessing,omitempty"`
	Amount               uint64   `protobuf:"varint,8,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountOne            uint64   `protobuf:"varint,9,opt,name=amount_one,json=amountOne,proto3" json:"amount_one,omitempty"`
	Quantity             uint32   `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty"`
	RemainAmount         uint64   `protobuf:"varint,11,opt,name=remain_amount,json=remainAmount,proto3" json:"remain_amount,omitempty"`
	RemainQuantity       uint32   `protobuf:"varint,12,opt,name=remain_quantity,json=remainQuantity,proto3" json:"remain_quantity,omitempty"`
	OrderType            uint32   `protobuf:"varint,13,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	Status               uint32   `protobuf:"varint,14,opt,name=status,proto3" json:"status,omitempty"`
	PayStatus            uint32   `protobuf:"varint,15,opt,name=pay_status,json=payStatus,proto3" json:"pay_status,omitempty"`
	ExpireAt             uint64   `protobuf:"varint,16,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	CreatedAt            uint64   `protobuf:"varint,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdateAt             uint64   `protobuf:"varint,18,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvelopeGoods) Reset()         { *m = EnvelopeGoods{} }
func (m *EnvelopeGoods) String() string { return proto.CompactTextString(m) }
func (*EnvelopeGoods) ProtoMessage()    {}
func (*EnvelopeGoods) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee266e8c558e9dc5, []int{0}
}

func (m *EnvelopeGoods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvelopeGoods.Unmarshal(m, b)
}
func (m *EnvelopeGoods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvelopeGoods.Marshal(b, m, deterministic)
}
func (m *EnvelopeGoods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvelopeGoods.Merge(m, src)
}
func (m *EnvelopeGoods) XXX_Size() int {
	return xxx_messageInfo_EnvelopeGoods.Size(m)
}
func (m *EnvelopeGoods) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvelopeGoods.DiscardUnknown(m)
}

var xxx_messageInfo_EnvelopeGoods proto.InternalMessageInfo

func (m *EnvelopeGoods) GetEnvelopeNo() string {
	if m != nil {
		return m.EnvelopeNo
	}
	return ""
}

func (m *EnvelopeGoods) GetEnvelopeType() uint32 {
	if m != nil {
		return m.EnvelopeType
	}
	return 0
}

func (m *EnvelopeGoods) GetAccountNo() string {
	if m != nil {
		return m.AccountNo
	}
	return ""
}

func (m *EnvelopeGoods) GetOriginEnvelopeNo() string {
	if m != nil {
		return m.OriginEnvelopeNo
	}
	return ""
}

func (m *EnvelopeGoods) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *EnvelopeGoods) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *EnvelopeGoods) GetBlessing() string {
	if m != nil {
		return m.Blessing
	}
	return ""
}

func (m *EnvelopeGoods) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *EnvelopeGoods) GetAmountOne() uint64 {
	if m != nil {
		return m.AmountOne
	}
	return 0
}

func (m *EnvelopeGoods) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *EnvelopeGoods) GetRemainAmount() uint64 {
	if m != nil {
		return m.RemainAmount
	}
	return 0
}

func (m *EnvelopeGoods) GetRemainQuantity() uint32 {
	if m != nil {
		return m.RemainQuantity
	}
	return 0
}

func (m *EnvelopeGoods) GetOrderType() uint32 {
	if m != nil {
		return m.OrderType
	}
	return 0
}

func (m *EnvelopeGoods) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *EnvelopeGoods) GetPayStatus() uint32 {
	if m != nil {
		return m.PayStatus
	}
	return 0
}

func (m *EnvelopeGoods) GetExpireAt() uint64 {
	if m != nil {
		return m.ExpireAt
	}
	return 0
}

func (m *EnvelopeGoods) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *EnvelopeGoods) GetUpdateAt() uint64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

type RedEnvelopeItem struct {
	ItemNo               string   `protobuf:"bytes,1,opt,name=item_no,json=itemNo,proto3" json:"item_no,omitempty"`
	EnvelopeNo           string   `protobuf:"bytes,2,opt,name=envelope_no,json=envelopeNo,proto3" json:"envelope_no,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserId               uint64   `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount               uint64   `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Quantity             uint64   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	RemainAmount         uint64   `protobuf:"varint,7,opt,name=remain_amount,json=remainAmount,proto3" json:"remain_amount,omitempty"`
	AccountNo            string   `protobuf:"bytes,8,opt,name=account_no,json=accountNo,proto3" json:"account_no,omitempty"`
	PayStatus            uint32   `protobuf:"varint,9,opt,name=pay_status,json=payStatus,proto3" json:"pay_status,omitempty"`
	IsLuckiest           bool     `protobuf:"varint,10,opt,name=is_luckiest,json=isLuckiest,proto3" json:"is_luckiest,omitempty"`
	Desc                 string   `protobuf:"bytes,11,opt,name=desc,proto3" json:"desc,omitempty"`
	CreatedAt            uint64   `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdateAt             uint64   `protobuf:"varint,13,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedEnvelopeItem) Reset()         { *m = RedEnvelopeItem{} }
func (m *RedEnvelopeItem) String() string { return proto.CompactTextString(m) }
func (*RedEnvelopeItem) ProtoMessage()    {}
func (*RedEnvelopeItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee266e8c558e9dc5, []int{1}
}

func (m *RedEnvelopeItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedEnvelopeItem.Unmarshal(m, b)
}
func (m *RedEnvelopeItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedEnvelopeItem.Marshal(b, m, deterministic)
}
func (m *RedEnvelopeItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedEnvelopeItem.Merge(m, src)
}
func (m *RedEnvelopeItem) XXX_Size() int {
	return xxx_messageInfo_RedEnvelopeItem.Size(m)
}
func (m *RedEnvelopeItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RedEnvelopeItem.DiscardUnknown(m)
}

var xxx_messageInfo_RedEnvelopeItem proto.InternalMessageInfo

func (m *RedEnvelopeItem) GetItemNo() string {
	if m != nil {
		return m.ItemNo
	}
	return ""
}

func (m *RedEnvelopeItem) GetEnvelopeNo() string {
	if m != nil {
		return m.EnvelopeNo
	}
	return ""
}

func (m *RedEnvelopeItem) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RedEnvelopeItem) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RedEnvelopeItem) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *RedEnvelopeItem) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *RedEnvelopeItem) GetRemainAmount() uint64 {
	if m != nil {
		return m.RemainAmount
	}
	return 0
}

func (m *RedEnvelopeItem) GetAccountNo() string {
	if m != nil {
		return m.AccountNo
	}
	return ""
}

func (m *RedEnvelopeItem) GetPayStatus() uint32 {
	if m != nil {
		return m.PayStatus
	}
	return 0
}

func (m *RedEnvelopeItem) GetIsLuckiest() bool {
	if m != nil {
		return m.IsLuckiest
	}
	return false
}

func (m *RedEnvelopeItem) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *RedEnvelopeItem) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *RedEnvelopeItem) GetUpdateAt() uint64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

func init() {
	proto.RegisterType((*EnvelopeGoods)(nil), "red.proto.envelope.EnvelopeGoods")
	proto.RegisterType((*RedEnvelopeItem)(nil), "red.proto.envelope.RedEnvelopeItem")
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor_ee266e8c558e9dc5) }

var fileDescriptor_ee266e8c558e9dc5 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0xdb, 0x3a,
	0x10, 0x84, 0x63, 0x45, 0x96, 0xd6, 0x96, 0x9d, 0xf0, 0x90, 0x47, 0xbc, 0x20, 0x88, 0x91, 0x1e,
	0x6a, 0xa0, 0x40, 0x2e, 0xfd, 0x02, 0x1f, 0x82, 0x22, 0x40, 0xe1, 0xa2, 0x6a, 0xef, 0x02, 0x23,
	0x2d, 0x02, 0xa2, 0x16, 0xa9, 0x92, 0x54, 0x51, 0x7d, 0x42, 0xbf, 0xab, 0x3f, 0x56, 0x70, 0x29,
	0xd9, 0x95, 0xdb, 0xe6, 0xc6, 0x9d, 0xd9, 0x1d, 0x91, 0x3b, 0x23, 0x58, 0xa2, 0xfa, 0x86, 0x7b,
	0xdd, 0xe0, 0x7d, 0x63, 0xb4, 0xd3, 0x8c, 0x19, 0xac, 0xc2, 0xf1, 0x7e, 0x60, 0xee, 0x7e, 0x46,
	0x90, 0x3d, 0xf4, 0xc5, 0x3b, 0xad, 0x2b, 0xcb, 0x6e, 0x61, 0x3e, 0xb0, 0x85, 0xd2, 0x7c, 0xb2,
	0x9e, 0x6c, 0xd2, 0x1c, 0x06, 0x68, 0xa7, 0xd9, 0x2b, 0xc8, 0x0e, 0x0d, 0xae, 0x6b, 0x90, 0x9f,
	0xad, 0x27, 0x9b, 0x2c, 0x5f, 0x0c, 0xe0, 0xe7, 0xae, 0x41, 0x76, 0x03, 0x20, 0xca, 0x52, 0xb7,
	0xca, 0x79, 0x91, 0x29, 0x89, 0xa4, 0x3d, 0xb2, 0xd3, 0xec, 0x0d, 0x5c, 0x6a, 0x23, 0x9f, 0xa5,
	0x2a, 0x8e, 0xc2, 0x3c, 0xa2, 0xae, 0x8b, 0x40, 0x3c, 0x1c, 0x3f, 0x78, 0x0d, 0x69, 0x6b, 0xd1,
	0x14, 0x4a, 0xd4, 0xc8, 0xcf, 0xa9, 0x29, 0xf1, 0xc0, 0x4e, 0xd4, 0xc8, 0xfe, 0x83, 0x19, 0x91,
	0xb2, 0xe2, 0xf1, 0x7a, 0xb2, 0x89, 0xf2, 0xd8, 0x97, 0x8f, 0x15, 0xfb, 0x1f, 0x92, 0xa7, 0x3d,
	0x5a, 0x2b, 0xd5, 0x33, 0x9f, 0x85, 0xa1, 0xa1, 0x66, 0x57, 0x10, 0x8b, 0xda, 0x5f, 0x85, 0x27,
	0x61, 0x26, 0x54, 0x74, 0x6b, 0x3a, 0x15, 0x5a, 0x21, 0x4f, 0x89, 0x4b, 0x03, 0xf2, 0x41, 0xa1,
	0x97, 0xfc, 0xda, 0x0a, 0xe5, 0xa4, 0xeb, 0x38, 0xd0, 0xa3, 0x0f, 0xb5, 0xdf, 0x8a, 0xc1, 0x5a,
	0x48, 0x55, 0xf4, 0xca, 0x73, 0x9a, 0x5e, 0x04, 0x70, 0x1b, 0xf4, 0x5f, 0xc3, 0xaa, 0x6f, 0x3a,
	0xe8, 0x2c, 0x48, 0x67, 0x19, 0xe0, 0x8f, 0x83, 0xda, 0x0d, 0x80, 0x36, 0x15, 0x9a, 0xb0, 0xe0,
	0x8c, 0x7a, 0x52, 0x42, 0x68, 0xbb, 0x57, 0x10, 0x5b, 0x27, 0x5c, 0x6b, 0xf9, 0x92, 0xa8, 0xbe,
	0xf2, 0x63, 0x8d, 0xe8, 0x8a, 0x9e, 0x5b, 0x85, 0xb1, 0x46, 0x74, 0x9f, 0x02, 0x7d, 0x0d, 0x29,
	0x7e, 0x6f, 0xa4, 0xc1, 0x42, 0x38, 0x7e, 0x41, 0xf7, 0x4b, 0x02, 0xb0, 0xa5, 0xb7, 0x97, 0x06,
	0x85, 0xc3, 0xca, 0xb3, 0x97, 0xe1, 0xed, 0x3d, 0xb2, 0x75, 0x64, 0x42, 0x53, 0x09, 0x47, 0xb3,
	0x2c, 0xcc, 0x06, 0x60, 0xeb, 0xee, 0x7e, 0x4c, 0x61, 0x95, 0x63, 0x35, 0x78, 0xf6, 0xe8, 0xb0,
	0xf6, 0xc6, 0x48, 0x87, 0xf5, 0x31, 0x43, 0xb1, 0x2f, 0x77, 0xfa, 0x34, 0x60, 0x67, 0x7f, 0x04,
	0x6c, 0xe4, 0xf7, 0xf4, 0xdf, 0x7e, 0x47, 0x23, 0xbf, 0x8f, 0x9e, 0x9e, 0x8f, 0x3c, 0xfd, 0xdd,
	0xb4, 0x90, 0x90, 0x17, 0x4c, 0x9b, 0xfd, 0xc5, 0xb4, 0x71, 0x94, 0x93, 0xd3, 0x28, 0x8f, 0x77,
	0x9e, 0x9e, 0xee, 0xfc, 0x16, 0xe6, 0xd2, 0x16, 0xfb, 0xb6, 0xfc, 0x22, 0xd1, 0x3a, 0x8a, 0x4d,
	0x92, 0x83, 0xb4, 0xef, 0x7b, 0x84, 0x31, 0x88, 0x2a, 0xb4, 0x25, 0xe5, 0x25, 0xcd, 0xe9, 0x7c,
	0xe2, 0xc5, 0xe2, 0x45, 0x2f, 0xb2, 0xb1, 0x17, 0x4f, 0x31, 0xfd, 0xe1, 0x6f, 0x7f, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x4b, 0x51, 0x0e, 0x5b, 0xfe, 0x03, 0x00, 0x00,
}
