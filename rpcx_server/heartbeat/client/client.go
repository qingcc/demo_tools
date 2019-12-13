package main

import (
	"context"
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "tcp@localhost:8006", "server address")
)


func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery(*addr, "")
	op := client.DefaultOption
	op.Heartbeat = true
	op.HeartbeatInterval = time.Second
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, op)
	defer xclient.Close()

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

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(time.Minute)
	}

}
