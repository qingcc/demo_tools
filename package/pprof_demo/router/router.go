package router

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//开启一个端口采集数据，可用于分析性能（cpu，内存， goroutine使用情况等等）
func InitPprof(addr string) {
	if addr == "" {
		addr = ":8081"
	}
	metricsRouter := http.DefaultServeMux
	metricsRouter.Handle("/metrics", promhttp.Handler())
	metricsRouter.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	adminServer := &http.Server{
		Addr:    addr,
		Handler: metricsRouter,
	}

	go func() {
		if err := adminServer.ListenAndServe(); err != nil {
			println("ListenAndServe metrics: ", err.Error())
		}
	}()
	select {}
}
