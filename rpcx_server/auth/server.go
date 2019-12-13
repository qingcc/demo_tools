package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"

	"github.com/qingcc/demo_tools/rpcx_server/example"
)

var (
	addr = flag.String("addr", "localhost:8008", "server address")

)

func main() {
	flag.Parse()
	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	s.AuthFunc = auth
	s.Serve("tcp", *addr)
	select {}
}

func auth(ctx context.Context, req *protocol.Message, token string) error {
	if token == "123456" {
		fmt.Println("token:", token)
		return nil
	}

	return errors.New("invalid token")
}