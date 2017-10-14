package main

import (
	"fmt"

	proto "github.com/gorpcserverpoc/proto"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// Greeter -
type Greeter struct{}

// Hello -
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name + "!"
	return nil
}

// Goodbye -
func (g *Greeter) Goodbye(ctx context.Context, req *proto.GoodbyeRequest, rsp *proto.GoodbyeResponse) error {
	rsp.Greeting = "Goodbye " + req.Name + "!"
	return nil
}

func main() {

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	handler := new(Greeter)
	service.Server().Handle(service.Server().NewHandler(handler))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
