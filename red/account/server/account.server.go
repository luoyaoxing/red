package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/account/logic"
	"gitlab.kay.com/red/proto/account"
	"log"
)

var (
	serverName = "AccountAo"
)

func init()  {
	err := logger.InitZapLogger(serverName)
	if err != nil {
		log.Fatal(err)
	}
}

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
		micro.Name("com.kay.red.AccountAo"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	service.Init()

	accountLogic, err := logic.NewAccountLogic()
	if err != nil {
		logger.Errorf("NewUserLogic err:%s", err.Error())
		return
	}

	red_proto_account.RegisterAccountAoServiceHandler(service.Server(), accountLogic)

	if err := service.Run(); err != nil {
		logger.Errorf("service Run err:%s", err.Error())
		return
	}
}
