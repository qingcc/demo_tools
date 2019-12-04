package main

import (
	"context"
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/smallnest/rpcx/client"
	"log"
)

var (
	zkAddr   = flag.String("zkAddr", "localhost:6033", "server address")
	basePath = flag.String("base", "rpcx_test", "base path")
)

func main() {
	flag.Parse()
	d := client.NewZookeeperDiscovery(*basePath, "Arith", []string{*zkAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	reply := &example.Reply{}
	args := example.Args{
		A: 10,
		B: 13,
	}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
