package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8009", nil))
	}()
	c := cron.New()
	c.AddFunc(fmt.Sprintf("0 */%d * * *", 5), func() {
		fmt.Println("aaa")
	})
	c.Start()
	select {}
}
