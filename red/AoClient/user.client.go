package AoClient

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gitlab.kay.com/config"
	"gitlab.kay.com/red/proto/user"
)

type UserAoClient struct {
	cli red_proto_user.UserAoService
}

func NewUserAoClient() *UserAoClient {
	userAoClient := new(UserAoClient)

	reg := consul.NewRegistry(func(options *registry.Options) {
		config.Get("consul", "host")
	})

	cli := client.NewClient(client.Registry(reg))
	userCli := red_proto_user.NewUserAoService("com.kay.red.UserAo", cli)

	userAoClient.cli = userCli
	return userAoClient
}
