package main

import (
	"context"
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "tcp@localhost:8020", "server address")
)

func main() {
	flag.Parse()

	//client_timeout()//client.DefaultOption 中修改配置参数
	client_context_timeout() //推荐用法
}

func client_timeout() {
	d := client.NewPeer2PeerDiscovery(*addr, "")
	//client端 设置超时限制
	op := client.DefaultOption
	op.ReadTimeout = time.Second * 10
	xclient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, op)
	defer xclient.Close()
	rpcx_server.Cal(xclient)
}

func client_context_timeout() {
	d := client.NewPeer2PeerDiscovery(*addr, "")
	//client端 设置超时限制

	xclient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()
	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &example.Reply{}
		ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
		err := xclient.Call(ctx, "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		cancelFn()
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(time.Second * 20)
	}
}
