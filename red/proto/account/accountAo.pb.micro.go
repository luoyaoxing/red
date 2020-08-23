// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: accountAo.proto

package red_proto_account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AccountAoService service

type AccountAoService interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...client.CallOption) (*CreateAccountResponse, error)
	Transfer(ctx context.Context, in *AccountTransferRequest, opts ...client.CallOption) (*AccountTransferResponse, error)
	StoreValue(ctx context.Context, in *StoreValueRequest, opts ...client.CallOption) (*StoreValueResponse, error)
	GetEnvelopeAccountByUserId(ctx context.Context, in *GetEnvelopeAccountByUserIdRequest, opts ...client.CallOption) (*GetEnvelopeAccountByUserIdResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*GetAccountResponse, error)
}

type accountAoService struct {
	c    client.Client
	name string
}

func NewAccountAoService(name string, c client.Client) AccountAoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "red.proto.account"
	}
	return &accountAoService{
		c:    c,
		name: name,
	}
}

func (c *accountAoService) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...client.CallOption) (*CreateAccountResponse, error) {
	req := c.c.NewRequest(c.name, "AccountAoService.CreateAccount", in)
	out := new(CreateAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAoService) Transfer(ctx context.Context, in *AccountTransferRequest, opts ...client.CallOption) (*AccountTransferResponse, error) {
	req := c.c.NewRequest(c.name, "AccountAoService.Transfer", in)
	out := new(AccountTransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAoService) StoreValue(ctx context.Context, in *StoreValueRequest, opts ...client.CallOption) (*StoreValueResponse, error) {
	req := c.c.NewRequest(c.name, "AccountAoService.StoreValue", in)
	out := new(StoreValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAoService) GetEnvelopeAccountByUserId(ctx context.Context, in *GetEnvelopeAccountByUserIdRequest, opts ...client.CallOption) (*GetEnvelopeAccountByUserIdResponse, error) {
	req := c.c.NewRequest(c.name, "AccountAoService.GetEnvelopeAccountByUserId", in)
	out := new(GetEnvelopeAccountByUserIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAoService) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*GetAccountResponse, error) {
	req := c.c.NewRequest(c.name, "AccountAoService.GetAccount", in)
	out := new(GetAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountAoService service

type AccountAoServiceHandler interface {
	CreateAccount(context.Context, *CreateAccountRequest, *CreateAccountResponse) error
	Transfer(context.Context, *AccountTransferRequest, *AccountTransferResponse) error
	StoreValue(context.Context, *StoreValueRequest, *StoreValueResponse) error
	GetEnvelopeAccountByUserId(context.Context, *GetEnvelopeAccountByUserIdRequest, *GetEnvelopeAccountByUserIdResponse) error
	GetAccount(context.Context, *GetAccountRequest, *GetAccountResponse) error
}

func RegisterAccountAoServiceHandler(s server.Server, hdlr AccountAoServiceHandler, opts ...server.HandlerOption) error {
	type accountAoService interface {
		CreateAccount(ctx context.Context, in *CreateAccountRequest, out *CreateAccountResponse) error
		Transfer(ctx context.Context, in *AccountTransferRequest, out *AccountTransferResponse) error
		StoreValue(ctx context.Context, in *StoreValueRequest, out *StoreValueResponse) error
		GetEnvelopeAccountByUserId(ctx context.Context, in *GetEnvelopeAccountByUserIdRequest, out *GetEnvelopeAccountByUserIdResponse) error
		GetAccount(ctx context.Context, in *GetAccountRequest, out *GetAccountResponse) error
	}
	type AccountAoService struct {
		accountAoService
	}
	h := &accountAoServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AccountAoService{h}, opts...))
}

type accountAoServiceHandler struct {
	AccountAoServiceHandler
}

func (h *accountAoServiceHandler) CreateAccount(ctx context.Context, in *CreateAccountRequest, out *CreateAccountResponse) error {
	return h.AccountAoServiceHandler.CreateAccount(ctx, in, out)
}

func (h *accountAoServiceHandler) Transfer(ctx context.Context, in *AccountTransferRequest, out *AccountTransferResponse) error {
	return h.AccountAoServiceHandler.Transfer(ctx, in, out)
}

func (h *accountAoServiceHandler) StoreValue(ctx context.Context, in *StoreValueRequest, out *StoreValueResponse) error {
	return h.AccountAoServiceHandler.StoreValue(ctx, in, out)
}

func (h *accountAoServiceHandler) GetEnvelopeAccountByUserId(ctx context.Context, in *GetEnvelopeAccountByUserIdRequest, out *GetEnvelopeAccountByUserIdResponse) error {
	return h.AccountAoServiceHandler.GetEnvelopeAccountByUserId(ctx, in, out)
}

func (h *accountAoServiceHandler) GetAccount(ctx context.Context, in *GetAccountRequest, out *GetAccountResponse) error {
	return h.AccountAoServiceHandler.GetAccount(ctx, in, out)
}