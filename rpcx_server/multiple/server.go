package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"log"
	"github.com/qingcc/demo_tools/rpcx_server/example"
)

var (
	addr1 = flag.String("addr1", ":8001", "server address")
	addr2 = flag.String("addr2", ":8002", "server address")
	addr3 = flag.String("addr3", ":8003", "server address")
)

func main()  {
	flag.Parse()
	go newServer(*addr1)
	go newServer(*addr2)
	go newServer(*addr3)

	select {}
}

func newServer(addr string)  {
	log.Printf("调用地址：%s", addr)
	s := server.NewServer()
	if err := s.RegisterName("Arith", new(example.Arith), ""); err != nil {
		log.Println(err)
	}

	if err := s.Serve("tcp", addr); err != nil {
		panic(err)
	}
}


