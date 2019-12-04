package main

import (
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server"
	"github.com/smallnest/rpcx/client"
	"time"
)

var (
	addr     = flag.String("addr", ":8000", "server address")
	addr1    = flag.String("addr1", ":8001", "server address")
	addr2    = flag.String("addr2", ":8002", "server address")
	group    = flag.String("g", "test", "group name")
	basePath = flag.String("basePath", "test", "base path")
	zkaddr   = flag.String("zkaddr", "localhost:6033", "zookeeper address")
)

func main() {
	flag.Parse()
	//peer2peer()//点对点
	//multiple()//点对多
	//zookeeper()//zookeeper plugin
	mdns() //mdns plugin
}

/*
TODO xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
TODO 失败模式（宕机）：FailMode的设置仅仅对同步调用有效(XClient.Call), 异步调用用，这个参数是无意义的。
TODO 	Failfast	//在这种模式下， 一旦调用一个节点失败， rpcx立即会返回错误。 注意这个错误不是业务上的 Error,
					业务上服务端返回的Error应该正常返回给客户端，这里的错误可能是网络错误或者服务异常。
		Failover	//在这种模式下, rpcx如果遇到错误，它会尝试调用另外一个节点， 直到服务节点能正常返回信息，
					或者达到最大的重试次数。 重试测试Retries在参数Option中设置， 缺省设置为3。
		Failtry		//在这种模式下， rpcx如果调用一个节点的服务出现错误， 它也会尝试，
					但是还是选择这个节点进行重试， 直到节点正常返回数据或者达到最大重试次数。
		Failbackup	//在这种模式下， 如果服务节点在一定的时间内不返回结果，
					rpcx客户端会发送相同的请求到另外一个节点， 只要这两个节点有一个返回， rpcx就算调用成功。
*/

/*
todo 点对点是最简单的一种注册中心的方式，事实上没有注册中心，客户端直接得到唯一的服务器的地址，连接服务。
	在系统扩展时，你可以进行一些更改，服务器不需要进行更多的配置 客户端使用Peer2PeerDiscovery来设置该服务的网络和地址。
	由于只有有一个节点，因此选择器是不可用的。
*/
func peer2peer() {
	d := client.NewPeer2PeerDiscovery(*addr, "group="+*group)
	//todo NewXClient必须使用服务名称作为第一个参数，然后使用failmode，selector，discovery和其他选项。
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()
	rpcx_server.Cal(xclient)
}

/*
todo 上面的方式只能访问一台服务器，假设我们有固定的几台服务器提供相同的服务，我们可以采用这种方式。
	如果你有多个服务但没有注册中心.你可以用编码的方式在客户端中配置服务的地址。 服务器不需要进行更多的配置。
	客户端使用MultipleServersDiscovery并仅设置该服务的网络和地址。
*/
func multiple() {
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()
	rpcx_server.Cal(xclient)
}

func zookeeper() {
	d := client.NewZookeeperDiscovery(*basePath, "Arith", []string{*zkaddr}, nil)
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()
	rpcx_server.Cal(xclient)
}

func mdns() {
	d := client.NewMDNSDiscovery("Arith", time.Second*10, time.Second*10, "")
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()
	rpcx_server.Cal(xclient)
}
