package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"log"
	"github.com/qingcc/demo_tools/rpcx_server/example"
)

var (
	addr = flag.String("addr", ":8977", "server address")
)

func main()  {
	flag.Parse()
	s := server.NewServer()
	if err := s.RegisterName("Arith", new(example.Arith), ""); err != nil {
		log.Println(err)
	}

	if err := s.Serve("tcp", *addr); err != nil {
		panic(err)
	}
}


