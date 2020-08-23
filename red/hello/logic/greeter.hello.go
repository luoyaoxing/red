package logic

import (
	"context"
	"gitlab.kay.com/red/proto/hello"
)

type Hello struct {}

func (h *Hello) Hello(ctx context.Context, req *red_proto_hello.HelloRequest, rsp *red_proto_hello.HelloResponse) error  {
	name := req.GetName()
	rsp.Greeting = "hello," + name
	return nil
}
