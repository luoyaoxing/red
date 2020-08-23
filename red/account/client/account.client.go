package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/proto/account"
)

func main() {
	err := config.LoadFile("/red/config/common.json")
	if err != nil {
		logger.Errorf("config LoadFile err:%s", err.Error())
		return
	}

	reg := consul.NewRegistry(func(options *registry.Options) {
		config.Get("consul", "host")
	})

	service := micro.NewService(
		micro.Name("com.kay.red.AccountAo.client"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	service.Init()

	accountCli := red_proto_account.NewAccountAoService("com.kay.red.AccountAo", service.Client())

	createAccount(accountCli)
}

func transfer(service red_proto_account.AccountAoService)  {
	body := &red_proto_account.TradeParticipator{
		AccountNo: "1gSFjdPpNlGq0YaFyL0CxB64NAW",
		AccountType: 1,
		UserId: 1,
		UserName: "daniel",
	}

	target := &red_proto_account.TradeParticipator{
		AccountNo: "1gSGGb83oEXvO5xiEYDrd7XrQ5A",
		AccountType: 1,
		UserId: 2,
		UserName: "Kay",
	}

	accountTransefer := &red_proto_account.AccountTransfer{
		TradeBody: body,
		TradeTarget: target,
		Amount: 500000,
		ChangeType: -1,
		ChangeFlag: -1,
		Desc: "daniel向Kay转五千块钱",
	}

	req := &red_proto_account.AccountTransferRequest{
		AccountTransfer: accountTransefer,
	}

	resp, err := service.Transfer(context.Background(), req)
	if err != nil {
		logger.Errorf("CreateUser err:%s", err.Error())
		return
	}

	fmt.Printf("TransferStatus:%d\n", resp.TransferStatus)
}

func storeValue(service red_proto_account.AccountAoService)  {
	body := &red_proto_account.TradeParticipator{
		AccountNo: "1gSFjdPpNlGq0YaFyL0CxB64NAW",
		AccountType: 1,
		UserId: 1,
		UserName: "daniel",
	}

	//target := &red_proto_account.TradeParticipator{
	//	AccountNo: "1gSFjdPpNlGq0YaFyL0CxB64NAW",
	//	AccountType: 1,
	//	UserId: 1,
	//	UserName: "daniel",
	//}

	accountTransefer := &red_proto_account.AccountTransfer{
		TradeBody: body,
		//TradeTarget: target,
		Amount: 0,
		ChangeType: 1,
		ChangeFlag: 1,
		Desc: "充值1000块钱",
	}

	req := &red_proto_account.StoreValueRequest{
		AccountTransfer: accountTransefer,
	}

	resp, err := service.StoreValue(context.Background(), req)
	if err != nil {
		logger.Errorf("CreateUser err:%s", err.Error())
		return
	}

	fmt.Printf("TransferStatus:%d\n", resp.TransferStatus)
}

func getAccount(service red_proto_account.AccountAoService)  {
	req := &red_proto_account.GetAccountRequest{
		AccountNo: "1gSFjdPpNlGq0YaFyL0CxB64NAW",
	}

	resp, err := service.GetAccount(context.Background(), req)
	if err != nil {
		logger.Errorf("CreateUser err:%s", err.Error())
		return
	}

	fmt.Printf("%v\n", resp.Account)
}

func getEnvelopeAccountByUserId(service red_proto_account.AccountAoService)  {
	req := &red_proto_account.GetEnvelopeAccountByUserIdRequest{
		UserId: 2,
	}

	resp, err := service.GetEnvelopeAccountByUserId(context.Background(), req)
	if err != nil {
		logger.Errorf("CreateUser err:%s", err.Error())
		return
	}

	fmt.Printf("%v\n", resp.Accounts)
}

func createAccount(service red_proto_account.AccountAoService)  {
	req := &red_proto_account.CreateAccountRequest{
		UserId: 3,
		UserName: "system",
		AccountName: "system的红包账户",
		AccountType: 1,
		CurrencyCode: "RMB",
		Amount: 100000000000,
	}

	resp, err := service.CreateAccount(context.Background(), req)
	if err != nil {
		logger.Errorf("CreateUser err:%s", err.Error())
		return
	}

	fmt.Printf("%v\n", resp.Account)
}
