package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "80", "listen address")

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bye bye ,this is v1 httpServer"))
	})
	log.Println("Starting v1 server ...")
	log.Println("listen address: ", *addr)
	log.Fatal(http.ListenAndServe(":"+*addr, nil))
}
