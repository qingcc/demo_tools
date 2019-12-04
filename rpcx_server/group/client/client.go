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
	addr1 = flag.String("addr1", ":8007", "server address")
	addr2 = flag.String("addr2", ":8008", "server address")
	group = flag.String("g", "test", "group name")
)

/*
client设置了group，就只能去对应的group去找服务
*/
func main() {
	flag.Parse()
	//d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key:*addr1}, {Key:*addr2, Value:"group=" + *group}})
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}})
	option := client.DefaultOption
	option.Group = *group
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, option)
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
		time.Sleep(1e9)
	}
}
