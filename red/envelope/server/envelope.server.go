package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/prometheus/common/log"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/envelope/logic"
	"gitlab.kay.com/red/proto/envelope"
)

var (
	serverName = "RedEnvelopeAo"
)

func init()  {
	err := logger.InitZapLogger(serverName)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	logger.Info("start %s", serverName)

	err := config.LoadFile("/red/config/common.json")
	if err != nil {
		logger.Errorf("config LoadFile err:%s", err.Error())
		log.Fatal(err)
	}

	reg := consul.NewRegistry(func(opts *registry.Options) {
		opts.Addrs = []string{
			config.Get("consul", "host"),
		}
	})

	service := micro.NewService(
		micro.Name("com.kay.red.RedEnvelopeAo"),
		micro.Version("latest"),
		micro.Registry(reg),
		)

	service.Init()

	handler, handlerErr := logic.NewRedEnvelopeLogic()
	if handlerErr != nil {
		logger.Errorf("%s NewRedEnvelopeHandler handlerErr:%s", handlerErr.Error())
		log.Fatal(handlerErr)
	}

	red_proto_envelope.RegisterRedEnvelopeAoHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		logger.Errorf("service run err:%s", err.Error())
		log.Fatal(err)
	}
}
