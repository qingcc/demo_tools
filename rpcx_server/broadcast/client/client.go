package main

import (
	"context"
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/smallnest/rpcx/client"
	"log"
	"runtime"
	"time"
)

var (
	addr1 = flag.String("addr1", "tcp@localhost:8008", "server address")
	addr2 = flag.String("addr2", "tcp@localhost:8009", "server address")
	saddr  = flag.String("saddr", "tcp@localhost:6035", "state service address")
)

func main() {
	flag.Parse()
	broadcast()
}
func broadcast() {
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}, {Key:*saddr}})
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()
	cal(xclient)
}

func cal(xclient client.XClient) {
	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &example.Reply{}
		err := xclient.Broadcast(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d, cal method:%s", args.A, args.B, reply.C, printCallerName())
		time.Sleep(1e9)
	}
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
