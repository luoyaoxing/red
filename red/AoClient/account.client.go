package AoClient

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gitlab.kay.com/config"
	"gitlab.kay.com/red/common"
	"gitlab.kay.com/red/proto/account"
	"gitlab.kay.com/red/proto/envelope"
)

type AccountAoClient struct {
	cli red_proto_account.AccountAoService
}

func NewAccountAoClient() *AccountAoClient {
	accountAoClient := new(AccountAoClient)

	reg := consul.NewRegistry(func(options *registry.Options) {
		config.Get("consul", "host")
	})

	accountCli := red_proto_account.NewAccountAoService("com.kay.red.AccountAo", client.NewClient(client.Registry(reg)))

	accountAoClient.cli = accountCli
	return accountAoClient
}

func (client *AccountAoClient) GetEnvelopeAccountByUserId(uid uint64) (*red_proto_account.Account, error)  {
	req := &red_proto_account.GetEnvelopeAccountByUserIdRequest{
		UserId: uid,
	}

	resp, err := client.cli.GetEnvelopeAccountByUserId(context.Background(), req)
	if err != nil {
		return nil, err
	}

	accounts := resp.Accounts
	return accounts[0], nil
}

func (client *AccountAoClient) RedSendOutTransfer(acc *red_proto_account.Account, goods *red_proto_envelope.EnvelopeGoods) int32 {
	body := &red_proto_account.TradeParticipator{
		AccountNo: acc.GetAccountNo(),
		AccountType: int32(acc.GetAccountType()),
		UserId: acc.GetUserId(),
		UserName: acc.GetUserName(),
	}

	// 转给系统用户
	target := &red_proto_account.TradeParticipator{
		AccountNo: "1gSKALde5FufrI8Za9fQiVj0psk",
		AccountType: 1,
		UserId: 3,
		UserName: "system",
	}

	accountTransefer := &red_proto_account.AccountTransfer{
		TradeBody: body,
		TradeTarget: target,
		Amount: goods.GetAmount(),
		ChangeType: common.FlagTransferOut,
		ChangeFlag: common.EnvelopeOutgoing,
		Desc: goods.GetBlessing(),
	}

	req := &red_proto_account.AccountTransferRequest{
		AccountTransfer: accountTransefer,
	}

	resp, _ := client.cli.Transfer(context.Background(), req)
	return resp.TransferStatus
}

func (client *AccountAoClient) RedReceiveTransfer(acc *red_proto_account.Account, item *red_proto_envelope.RedEnvelopeItem) int32 {
	body := &red_proto_account.TradeParticipator{
		AccountNo: acc.GetAccountNo(),
		AccountType: int32(acc.GetAccountType()),
		UserId: acc.GetUserId(),
		UserName: acc.GetUserName(),
	}

	// 转给系统用户
	target := &red_proto_account.TradeParticipator{
		AccountNo: "1gSKALde5FufrI8Za9fQiVj0psk",
		AccountType: 1,
		UserId: 3,
		UserName: "system",
	}

	accountTransefer := &red_proto_account.AccountTransfer{
		TradeBody: body,
		TradeTarget: target,
		Amount: item.GetAmount(),
		ChangeType: common.FlagTransferIn,
		ChangeFlag: common.EnvelopeIncoming,
		Desc: item.GetDesc(),
	}

	req := &red_proto_account.AccountTransferRequest{
		AccountTransfer: accountTransefer,
	}

	resp, _ := client.cli.Transfer(context.Background(), req)
	return resp.TransferStatus
}