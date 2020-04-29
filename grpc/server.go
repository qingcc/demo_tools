package main

import (
	"context"
	"fmt"
	"github.com/qingcc/demo_tools/grpc/helloService"
	"net"

	"google.golang.org/grpc"
)

type HelloServiceServerImpl struct{}

func (s *HelloServiceServerImpl) SayHello(c context.Context, req *helloService.Request) (*helloService.Response, error) {
	fmt.Printf("%s\n", string(req.Data))

	resp := helloService.Response{}
	resp.Data = []byte("hello from server")

	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:57501")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	helloService.RegisterHelloServiceServer(s, &HelloServiceServerImpl{})
	fmt.Printf("Server listening on 127.0.0.1:57501\n")
	s.Serve(lis)
}
