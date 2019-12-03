package main

import (
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

var (
	addr = flag.String("addr", ":8006", "server address")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	addRegisterPlugin(s)
	if err := s.RegisterName("Arith", new(example.Arith), ""); err != nil {
		log.Println(err)
	}
	if err := s.Serve("tcp", *addr); err != nil {
		log.Printf(err.Error())
	}
}

func addRegisterPlugin(s *server.Server) {
	r := serverplugin.NewMDNSRegisterPlugin("tcp@"+*addr, 0, metrics.NewRegistry(), time.Minute, "")

	if err := r.Start(); err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
