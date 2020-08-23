package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/proto/user"
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
		micro.Name("com.kay.red.UserAo.client"),
		micro.Version("latest"),
		micro.Registry(reg),
		)

	service.Init()

	userCli := red_proto_user.NewUserAoService("com.kay.red.UserAo", service.Client())

	createUser(userCli)

	//getUser(userCli)
}

func getUser(service red_proto_user.UserAoService)  {
	req := &red_proto_user.GetUserRequest{
		UserId: 1,
	}

	resp, err := service.GetUser(context.Background(), req)
	if err != nil {
		logger.Errorf("CreateUser err:%s", err.Error())
		return
	}

	fmt.Printf("%v\n", resp.User)
}

func createUser(service red_proto_user.UserAoService)  {
	req := &red_proto_user.CreateUserRequest{
		UserName: "system",
		NickName: "system",
	}

	resp, err := service.CreateUser(context.Background(), req)
	if err != nil {
		logger.Errorf("CreateUser err:%s", err.Error())
		return
	}

	fmt.Println("用户uid:", resp.UserId)
}
