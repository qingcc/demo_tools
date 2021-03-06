package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/nacos-group/nacos-sdk-go/utils"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestGetConn(t *testing.T) {
	c := GetConn()
	defer c.Close()

	//region Remark: kv operation Author:qing
	//set timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	//1.add key value
	p, err := c.Put(ctx, "demo/add_key", "val1")
	if err != nil {
		fmt.Println("put into etcd failed,err:", err)
	}
	p, err = c.Put(ctx, "demo/add_key1", "add_key_val1")
	if err != nil {
		fmt.Println("put into etcd failed,err:", err)
	}
	p, err = c.Put(ctx, "demo/add_key2", "add_key_val2")
	if err != nil {
		fmt.Println("put into etcd failed,err:", err)
	}
	fmt.Printf("put result:%s", utils.ToJsonString(p))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	cancel()

	//2. get value by key
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	getResp, err := c.Get(ctx, "demo/add_key")
	if err != nil {
		fmt.Println("get from etcd failed,err:", err)
	}
	cancel()
	for _, item := range getResp.Kvs {
		fmt.Printf("get: %s : %s, version:%d", item.Key, item.Value, item.Version)
		fmt.Println()
	}

	//2.1. 通过prefix获取
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	// Get查询还可以增加WithPrefix选项，获取某个目录下的所有子元素
	resPre, err := c.Get(ctx, "demo/", clientv3.WithPrefix())
	if resPre != nil && err == nil {
		for _, item := range resPre.Kvs {
			fmt.Printf("get by prev: %s : %s, version:%d", item.Key, item.Value, item.Version)
			fmt.Println()
		}
	} else {
		log.Println("err:", err, "resPre:", utils.ToJsonString(resPre))
	}
	cancel()

	//3. update
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := c.Put(ctx, "demo/add_key", "update_value", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println("update etcd failed,err:", err)
	}
	fmt.Printf("update result:%s", string(resp.PrevKv.Value))
	fmt.Println()
	cancel()

	//4. delete
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	del, err := c.Delete(ctx, "demo/add_key")
	if err != nil {
		fmt.Println("delete etcd failed,err:", err)
	}
	fmt.Printf("delete result:%s, resp:%s", utils.ToJsonString(del), del.PrevKvs)
	cancel()
	//endregion
}

//watch
func TestWatch(t *testing.T) {
	c := GetConn()

	defer c.Close()

	c.Put(context.Background(), "demo/demo_watch", "watch_value")

	go func() {
		watchKey := c.Watch(context.Background(), "demo/demo_watch")
		for resp := range watchKey {
			for _, ev := range resp.Events {
				if ev != nil && ev.Kv != nil {
					fmt.Println(string(ev.Kv.Key), ":", string(ev.Kv.Value)) //监听 demo/demo_watch 中存储的值是否变化
				} else {
					log.Println("ev:", utils.ToJsonString(ev))
				}

			}
		}
	}()

	for i := 0; i < 10; i++ {
		val := "watch_value_changed_" + strconv.Itoa(i)
		_, err := c.Put(context.Background(), "demo/demo_watch", val) //不断向 demo/demo_watch 更新新的值
		if err != nil {
			fmt.Println("err:", err)
		}
		time.Sleep(time.Second)
	}
}

func TestTransaction(t *testing.T) {
	c := GetConn()

	var wg sync.WaitGroup
	wg.Add(10)

	key10 := "setnx"
	for i := 0; i < 10; i++ {
		go func(i int) {
			t, err := c.Txn(context.Background()).
				If(clientv3.Compare(clientv3.CreateRevision(key10), "=", 0)).
				Then(clientv3.OpPut(key10, fmt.Sprintf("%d", i))).
				Commit()

			if err != nil {
				log.Println("err:", err)
			}
			log.Println("txn:", utils.ToJsonString(t))
			wg.Done()
		}(i)
	}
	wg.Wait()
	if res, err := c.Get(context.TODO(), key10); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("last get : ", utils.ToJsonString(res))
	}
}

func TestLease(t *testing.T) {
	c := GetConn()

	kv := clientv3.NewKV(c)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	key := "demo/demo1_key"
	//delete all keys
	kv.Delete(ctx, key, clientv3.WithPrefix())

	l := clientv3.NewLease(c)

	//lease
	lease, err := c.Grant(ctx, 3)
	if err != nil {
		panic("grant lease failed, err:" + err.Error())
	}
	//insert key with a lease of 3 second ttl
	kv.Put(ctx, key, "release_value_of_3_second", clientv3.WithLease(lease.ID))

	r, _ := kv.Get(ctx, key)
	if len(r.Kvs) == 0 {
		log.Println("no more key")
	} else {
		for _, v := range r.Kvs {
			fmt.Println(fmt.Sprintf("%s:%s", v.Key, string(v.Value)))
		}
	}

	time.Sleep(time.Second * 2)
	fmt.Println("after sleep 2 second")
	r, _ = kv.Get(ctx, key)
	if len(r.Kvs) == 0 {
		log.Println("no more key")
	} else {
		for _, v := range r.Kvs {
			fmt.Println(fmt.Sprintf("%s:%s", v.Key, string(v.Value)))
		}
	}
	time.Sleep(time.Second * 2)

	fmt.Println("after sleep 3 second")
	r, _ = kv.Get(ctx, key)
	if len(r.Kvs) == 0 {
		log.Println("no more key")
	} else {
		for _, v := range r.Kvs {
			fmt.Println(fmt.Sprintf("%s:%s", v.Key, string(v.Value)))
		}
	}

	//续租
	respLease, _ := c.Grant(ctx, 2)
	ctx1, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	respLeaseChan, err := c.KeepAlive(ctx1, respLease.ID)

	//watch lease
	go func() {
		for {
			select {
			case res := <-respLeaseChan:
				if res == nil {
					fmt.Println("已经关闭续租功能")
					return
				} else {
					fmt.Println("续租成功")
					goto END
				}
			}
		}
	END:
		time.Sleep(time.Microsecond * 500)
	}()

	//watch lease
	go func() {
		wchan := c.Watch(ctx, key)
		for w := range wchan {
			for _, e := range w.Events {
				fmt.Println(fmt.Sprintf("watch version: %d ==============%s:%s", e.Kv.Version, e.Kv.Key, string(e.Kv.Value)))
			}
		}
	}()

	//lease

	//insert key with a lease of 3 second ttl
	kv.Put(ctx, key, "release_value_of_3_second", clientv3.WithLease(respLease.ID))

	cancel()
	time.Sleep(time.Second * 5)

	_, err = l.Revoke(ctx, lease.ID)
	if err != nil {
		panic("撤销租约失败" + err.Error())
	}
	fmt.Println("撤销租约成功")

	r, _ = kv.Get(ctx, key)
	if len(r.Kvs) == 0 {
		log.Println("no more key")
	} else {
		for _, v := range r.Kvs {
			fmt.Println(fmt.Sprintf("%s:%s", v.Key, string(v.Value)))
		}
	}
	time.Sleep(time.Second * 20)
}
