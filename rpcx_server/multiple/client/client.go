package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8977", "server address")
	kv = []*client.KVPair{{Key:":8001"},{Key:":8002"},{Key:":8003"}}
)

func main()  {
	flag.Parse()

	for i:=0; ; i++ {
		cli(i)
		time.Sleep(time.Second*1)
		//if i % 10 == 0 && i != 0 && i <31 {
		//	cl := client.MultipleServersDiscovery{}
		//	cl.Update(kv[i%10:])
		//}
	}
}

func cli(i int)  {
	d := client.NewMultipleServersDiscovery(kv)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	reply := &example.Reply{}
	args := example.Args{
		A: 10,
		B: i,
	}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
