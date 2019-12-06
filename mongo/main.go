package main

import (
	"github.com/qingcc/demo_tools/mongo/demo"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//add()

	demo.MulQuery()
	select {}
}

var (
	skill = []string{"Go project", "mongodb", "redis", "docker", "rpcx", "rabbitmq", "zookeeper", "shell", "git", "http", "tcp", "linux"}
)

func add() {
	go func() {
		i := 0
		for {
			item := demo.Books{
				Name:   "mongo",
				Price:  rand.Float64(),
				Author: "author_" + strconv.Itoa(int(time.Now().Unix())),
				AuthorInfo: []demo.AuthorInfo{
					{
						Username:      "zhangsan",
						Age:           rand.Intn(50),
						Phone:         "15212341" + strconv.Itoa(int(time.Now().Unix()))[:3],
						Favoritebooks: skill[:rand.Intn(11)],
					},
				},
			}
			i++
			demo.InsertBooks(item)
			if i > 10 {
				log.Println("---------end---------")
				time.Sleep(time.Hour)
			}
		}
	}()
}
