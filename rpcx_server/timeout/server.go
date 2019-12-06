package main

import (
	"context"
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/smallnest/rpcx/server"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8020", "server address")
)

func main() {
	flag.Parse()
	//server端 设置超时限制
	//s := server.NewServer(server.WithReadTimeout(time.Second*10), server.WithWriteTimeout(time.Second*20))
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	time.Sleep(12 * time.Second)
	reply.C = args.A * args.B
	return nil
}
