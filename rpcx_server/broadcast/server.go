package main

import (
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server"
)

var (
	addr1 = flag.String("addr1", "localhost:8008", "server address")
	addr2 = flag.String("addr2", "localhost:8009", "server address")
	saddr  = flag.String("saddr", "localhost:6035", "state service address")
)

func main() {
	flag.Parse()
	go rpcx_server.CreateServer(*addr1, "", "")
	//go rpcx_server.CreateServer(*addr2, "", "1")
	go state(*saddr, "state=inactive")
	select {}
}


func state(addr, meta string)  {
	rpcx_server.CreateServer(addr, meta, "1")
}