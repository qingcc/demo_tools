package main

import (
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server"
	"github.com/qingcc/demo_tools/rpcx_server/example"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

var (
	addr   = flag.String("addr", ":8000", "server address")
	addr1  = flag.String("addr1", ":8001", "server address")
	addr2  = flag.String("addr2", ":8002", "server address")
	zkaddr = flag.String("zkaddr", ":6000", "server address")
	zkser  = flag.String("zkser", "localhost:6033", "zookeeper address")
	maddr  = flag.String("maddr", ":6034", "mdns service address")

	basePath = flag.String("basePath", "test", "base path")
	group    = flag.String("g", "test", "group name")
)

func main() {
	flag.Parse()
	go rpcx_server.CreateServer(*addr, "", "")
	go rpcx_server.CreateServer(*addr1, "", "1")
	go rpcx_server.CreateServer(*addr2, *group, "")
	go zk(*zkaddr, "")
	go mdns(*maddr, "")
	select {}

}

func zk(addr, meta string) {
	s := server.NewServer()
	addzk(s)
	s.RegisterName("Arith", new(example.Arith), meta)
	s.Serve("tcp", addr)
}

func addzk(s *server.Server) {
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *zkaddr,
		ZooKeeperServers: []string{*zkser},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}

	if err := r.Start(); err != nil {
		log.Println("zookeeper plugin service failed:", err.Error())
	}

	s.Plugins.Add(r)
}

func mdns(addr, meta string) {
	s := server.NewServer()
	addMdns(s)
	s.RegisterName("Arith", new(example.Arith), meta)
	s.Serve("tcp", addr)
}

func addMdns(s *server.Server) {
	r := serverplugin.NewMDNSRegisterPlugin("tcp@"+*maddr, 8972, metrics.NewRegistry(), time.Minute, "")
	if err := r.Start(); err != nil {
		log.Println("mdns server failed:", err.Error())
	}
	s.Plugins.Add(r)

}
