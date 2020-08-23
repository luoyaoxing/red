package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/prometheus/common/log"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/hello/logic"
	"gitlab.kay.com/red/proto/hello"
)

var (
	serverName = "hello"
)

func init()  {
	err := logger.InitZapLogger(serverName)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	logger.Info("com.kay.red.hello start")


	err := config.LoadFile("/red/hello/config/common.json")
	if err != nil {
		logger.Errorf("config LoadFile err:%s", err.Error())
		return
	}

	logger.Infof("config map:%#v", config.Map())
	logger.Infof("config get:%s", config.Get("mysql", "host"))


	reg := consul.NewRegistry(func(opts *registry.Options) {
		opts.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("com.kay.red.hello"),
		micro.Version("latest"),
		micro.Registry(reg),
		)

	service.Init()

	red_proto_hello.RegisterGreeterHandler(service.Server(), new(logic.Hello))

	if err := service.Run(); err != nil {
		logger.Errorf("service run err:%s", err.Error())
		log.Fatal(err)
	}
}
