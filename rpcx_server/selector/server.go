package main

import (
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server"
)

var (
	addr1 = flag.String("addr1", "localhost:8010", "server address")
	addr2 = flag.String("addr2", "localhost:8011", "server address")
)

func main() {
	flag.Parse()
	go rpcx_server.CreateServer(*addr1, "weight=7", "")
	go rpcx_server.CreateServer(*addr2, "weight=3", "1")
	select {}
}
