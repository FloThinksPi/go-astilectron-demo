package main

import (
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/transport"

	hello "github.com/FloThinksPi/go-harmony/protocol/hello"

	"context"
)

func init() {
	client.DefaultClient.Init(
		client.Transport(
			transport.NewTransport(transport.Secure(true)),
		),
	)
}

func main() {

	cl := hello.SayServiceClient("go.micro.srv.greeter", client.DefaultClient)

	rsp, err := cl.Hello(context.TODO(), &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
