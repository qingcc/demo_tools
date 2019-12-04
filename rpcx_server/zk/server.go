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
	addr     = flag.String("addr", ":8005", "server address")
	zkaddr   = flag.String("zkaddr", "localhost:6033", "zookeeper address")
	basePath = flag.String("basePath", "rpcx_test", "base path")
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
	//todo & 取地址 一定要取地址，否则报错
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{*zkaddr},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}

	if err := r.Start(); err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
