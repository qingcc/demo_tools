package main

import (
	"context"
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
	"log"
	"time"
)

var(
	addr = flag.String("addr", "tcp@localhost:8008", "server address")
)

func main()  {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery(*addr, "")
	op := client.DefaultOption
	op.ReadTimeout = time.Second*10
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, op)
	defer xclient.Close()

	xclient.Auth("1231456")

	args := &example.Args{
		A: 2,
		B: 3,
	}
	reply := &example.Reply{}
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, make(map[string]string))
	err := xclient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)


}
