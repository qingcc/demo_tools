package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"log"
	"github.com/qingcc/demo_tools/rpcx_server/example"
)

var (
	addr = flag.String("addr", "localhost:8977", "server address")
)

func main()  {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	reply := &example.Reply{}
	args := example.Args{
		A: 10,
		B: 13,
	}
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	}else {
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
}

