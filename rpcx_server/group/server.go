package main

import (
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/smallnest/rpcx/server"
	"log"
)

var (
	addr1 = flag.String("addr1", ":8007", "server address")
	addr2 = flag.String("addr2", ":8008", "server address")
)

func main() {
	flag.Parse()
	go create(*addr1, "")
	go create(*addr2, "group=test")
	select {}

}

func create(addr, meta string) {
	s := server.NewServer()

	if err := s.RegisterName("Arith", new(example.Arith), meta); err != nil {
		log.Printf("RegisterName failed: %s", err.Error())
	}
	if err := s.Serve("tcp", addr); err != nil {
		log.Println(err.Error())
	}
}
