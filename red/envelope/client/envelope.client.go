package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/proto/envelope"
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
		micro.Name("com.kay.red.RedEnvelopeAo.client"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	service.Init()

	envelopeCli := red_proto_envelope.NewRedEnvelopeAoService("com.kay.red.RedEnvelopeAo", service.Client())

	listItems(envelopeCli)
}

func listItems(service red_proto_envelope.RedEnvelopeAoService)  {
	req := &red_proto_envelope.ListItemsRequest{
		EnvelopeNo: "1gUmWb0fTKJwmkUCNTJXSqHX4M4",
	}

	resp, err := service.ListItems(context.Background(), req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Printf("RedEnvelopeItem:%v\n", resp.EnvelopeItems)
}

func listReceivable(service red_proto_envelope.RedEnvelopeAoService)  {
	req := &red_proto_envelope.ListReceivableRequest{
		Page: 2,
		Size: 2,
	}

	resp, err := service.ListReceivable(context.Background(), req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Printf("RedEnvelopeItem:%v\n", resp.EnvelopeGoods)
	fmt.Printf("totalPage:%v\n", resp.TotalPage)
}

func listReceived(service red_proto_envelope.RedEnvelopeAoService)  {
	req := &red_proto_envelope.ListReceivedRequest{
		UserId: 2,
	}

	resp, err := service.ListReceived(context.Background(), req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Printf("RedEnvelopeItem:%v\n", resp.EnvelopeItems)
}

func listSent(service red_proto_envelope.RedEnvelopeAoService)  {
	req := &red_proto_envelope.ListSentRequest{
		UserId: 2,
	}

	resp, err := service.ListSent(context.Background(), req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Printf("RedEnvelopeItem:%v\n", resp.EnvelopeGoods)
}

func get(service red_proto_envelope.RedEnvelopeAoService)  {
	req := &red_proto_envelope.GetRedEnvelopRequest{
		EnvelopeNo: "1gUmWb0fTKJwmkUCNTJXSqHX4M4",
	}

	resp, err := service.Get(context.Background(), req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Printf("RedEnvelopeItem:%v\n", resp.EnvelopeGoods)
}

func receiver(service red_proto_envelope.RedEnvelopeAoService)  {
	req := &red_proto_envelope.RedEnvelopeReceiveRequest{
		EnvelopeNo: "1gUmWb0fTKJwmkUCNTJXSqHX4M4",
		UserId: 2,
		UserName: "Kay",
		AccountNo: "1gSGGb83oEXvO5xiEYDrd7XrQ5A",
	}

	resp, err := service.Receive(context.Background(), req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Printf("RedEnvelopeItem:%v\n", resp.RedEnvelopeItem)
}

func sendOut(service red_proto_envelope.RedEnvelopeAoService)  {
	req := &red_proto_envelope.RedEnvelopeSendRequest{
		EnvelopeType: 2,
		UserId: 1,
		UserName: "daniel",
		Amount: 200000,
		Quantity: 10,
	}

	resp, err := service.SendOut(context.Background(), req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Printf("EnvelopeGoods:%v\n", resp.EnvelopeGoods)
	fmt.Printf("link:%v\n", resp.Link)
}
