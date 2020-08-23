package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/prometheus/common/log"
	"gitlab.kay.com/red/proto/hello"
)

func main() {
	reg := consul.NewRegistry(func(opts *registry.Options) {
		opts.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("com.kay.red.hello.client"),
		micro.Version("latest"),
		micro.Registry(reg),
		)

	service.Init()

	helloCli := red_proto_hello.NewGreeterService("com.kay.red.hello", service.Client())
	rsp, err := helloCli.Hello(context.Background(), &red_proto_hello.HelloRequest{Name: "kay"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp.Greeting)
}
