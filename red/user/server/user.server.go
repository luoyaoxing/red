package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/proto/user"
	"gitlab.kay.com/red/user/logic"
	"log"
)

var (
	serverName = "UserAo"
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
		micro.Name("com.kay.red.UserAo"),
		micro.Version("latest"),
		micro.Registry(reg),
		)

	service.Init()

	userLogic, err := logic.NewUserLogic()
	if err != nil {
		logger.Errorf("NewUserLogic err:%s", err.Error())
		return
	}

	red_proto_user.RegisterUserAoServiceHandler(service.Server(), userLogic)

	if err := service.Run(); err != nil {
		logger.Errorf("service Run err:%s", err.Error())
		return
	}
}
