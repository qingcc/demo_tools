package etcd

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"sync"
	"time"
)

var (
	etcdConn *clientv3.Client
	etcdOnce sync.Once
)

func GetConn() *clientv3.Client {
	etcdOnce.Do(func() {
		etcdConn = getConn()
	})
	return etcdConn
}

func getConn() *clientv3.Client {
	//客户端配置
	config := clientv3.Config{
		Endpoints:   []string{"47.112.210.86:7379"},
		DialTimeout: 5 * time.Second,
	}

	//建立连接
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println("err:", err)
		return etcdConn
	}
	fmt.Println("connect to etcd success")
	etcdConn = client
	return etcdConn
}
