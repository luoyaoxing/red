// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package red_proto_user

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

type RequestHeader struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Chn                  string   `protobuf:"bytes,2,opt,name=chn,proto3" json:"chn,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RequestHeader) GetChn() string {
	if m != nil {
		return m.Chn
	}
	return ""
}

func (m *RequestHeader) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ResponseHeader struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseHeader) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "red.proto.user.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "red.proto.user.ResponseHeader")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x86, 0x83, 0x18, 0x8c, 0x17, 0x21, 0xe6, 0x06, 0xd3, 0xd1, 0x30, 0x39, 0xb1, 0x98, 0xf8,
	0x0c, 0xce, 0xf5, 0x09, 0xb4, 0xbd, 0xa0, 0x43, 0x7b, 0x78, 0x47, 0x7d, 0x7e, 0x43, 0x81, 0xed,
	0xfb, 0xbf, 0xe1, 0xcb, 0x0f, 0x07, 0xc7, 0x21, 0x70, 0xec, 0x06, 0xe1, 0x91, 0xb1, 0x11, 0xf2,
	0x33, 0x76, 0x49, 0x49, 0xda, 0x07, 0xd4, 0x96, 0xbe, 0x89, 0x74, 0xbc, 0xd3, 0xd3, 0x93, 0xe0,
	0x09, 0x2a, 0xe5, 0x24, 0x8e, 0x4c, 0x71, 0x2e, 0x2e, 0x7b, 0xbb, 0x2c, 0x3c, 0x42, 0xe9, 0xde,
	0xd1, 0x6c, 0xb2, 0x9c, 0x10, 0x0d, 0xec, 0x7e, 0x24, 0xfa, 0xe1, 0x68, 0xca, 0x6c, 0xd7, 0xd9,
	0xde, 0xa0, 0xb1, 0xa4, 0x03, 0x47, 0xa5, 0xa5, 0x8a, 0xb0, 0x75, 0xec, 0xe7, 0x66, 0x6d, 0x33,
	0x4f, 0xc5, 0xa0, 0xfd, 0x5a, 0x0c, 0xda, 0xbf, 0xaa, 0x7c, 0xec, 0xfa, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x14, 0x33, 0x4d, 0x26, 0xb3, 0x00, 0x00, 0x00,
}