package rpcx_server

import (
	"context"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"log"
	"runtime"
	"time"
)

func Cal(xclient client.XClient) {
	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d, cal method:%s", args.A, args.B, reply.C, printCallerName())
		time.Sleep(1e9)
	}
}

func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func CreateServer(addr, meta, t string) {
	s := server.NewServer()
	if t == "" {
		s.RegisterName("Arith", new(example.Arith), meta)
	} else {
		s.RegisterName("Arith", new(example.Arith1), meta)
	}

	s.Serve("tcp", addr)
}
