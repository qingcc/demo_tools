package main

import (
	"flag"
	"github.com/qingcc/demo_tools/rpcx_server"
	"github.com/smallnest/rpcx/client"
)

/*
todo 在大型的微服务系统中，我们会为同一个服务部署多个节点， 以便服务可以支持大并发的访问。
	它们可能部署在同一个数据中心的多个节点，或者多个数据中心中。
	那么， 客户端该如何选择一个节点呢？ rpcx通过 Selector来实现路由选择， 它就像一个负载均衡器，帮助你选择出一个合适的节点。
	rpcx提供了多个路由策略算法，你可以在创建XClient来指定。
	注意，这里的路由是针对 ServicePath 和 ServiceMethod的路由。
*/
var (
	addr1 = flag.String("addr1", "tcp@localhost:8010", "server address")
	addr2 = flag.String("addr2", "tcp@localhost:8011", "server address")

	xclient client.XClient
)

func main() {
	flag.Parse()
	randomSelect()
	//roundRobin()
	//weightedRoundRobin()
	//weightedICMP()
	//consistentHash()
	//geo()
	defer xclient.Close()
	rpcx_server.Cal(xclient)
}

/*
todo (random)RandomSelect		随机	从配置的节点中随机选择一个节点。
			最简单，但是有时候单个节点的负载比较重。这是因为随机数只能保证在大量的请求下路由的比较均匀，
			并不能保证在很短的时间内负载是均匀的。
*/
func randomSelect() {
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient = client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	return
}

/*
todo (roundrobin)RoundRobin			轮询	使用轮询的方式，依次调用节点，能保证每个节点都均匀的被访问。在节点的服务能力都差不多的时候适用。
*/

func roundRobin() {
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient = client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	return
}

/*
todo (weighted)WeightedRoundRobin	权重	使用Nginx 平滑的基于权重的轮询算法。
			比如如果三个节点a、b、c的权重是{ 5, 1, 1 }, 这个算法的调用顺序是 { a, a, b, a, c, a, a }, 相比较
			{ c, b, a, a, a, a, a }, 虽然权重都一样，但是前者更好，不至于在一段时间内将请求都发送给a。
*/
func weightedRoundRobin() {
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient = client.NewXClient("Arith", client.Failtry, client.WeightedRoundRobin, d, client.DefaultOption)
	return
}

/*
TODO (ping)WeightedICMP		网络质量优先		首先客户端会基于ping(ICMP)探测各个节点的网络质量，越短的ping时间，这个节点的权重也就越高。但是，我们也会保证网络较差的节点也有被调用的机会。
			假定t是ping的返回时间， 如果超过1秒基本就没有调用机会了:
			weight=191 if t <= 10
			weight=201 -t if 10 < t <=200
			weight=1 if 200 < t < 1000
			weight=0 if t >= 1000
*/
func weightedICMP() {
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient = client.NewXClient("Arith", client.Failtry, client.WeightedICMP, d, client.DefaultOption)
	return
}

/*
todo (hash)ConsistentHash		一致性哈希	使用 JumpConsistentHash 选择节点， 相同的servicePath, serviceMethod 和 参数会路由到同一个节点上。
			JumpConsistentHash 是一个快速计算一致性哈希的算法，但是有一个缺陷是它不能删除节点，如果删除节点，路由就不准确了，
			所以在节点有变动的时候它会重新计算一致性哈希。
*/
func consistentHash() {
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1, Value: ""}, {Key: *addr2}})
	xclient = client.NewXClient("Arith", client.Failtry, client.ConsistentHash, d, client.DefaultOption)
	return
}

/*
todo (geo)ConsistentHash		地理位置优先		它要求服务在注册的时候要设置它所在的地理经纬度。
			如果两个服务的节点的经纬度是一样的， rpcx会随机选择一个。
			比必须使用下面的方法配置客户端的经纬度信息：
			func (c *xClient) ConfigGeoSelector(latitude, longitude float64)
*/
func geo() {
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1, Value: "latitude=39.9289&longitude=116.3883"},
		{Key: *addr2, Value: "latitude=139.3453&longitude=23.3243"}})
	xclient = client.NewXClient("Arith", client.Failtry, client.ConsistentHash, d, client.DefaultOption)
	return
}
