package main

import (
	"context"
	"fmt"
	"github.com/qingcc/demo_tools/grpc/helloService"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:57501", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	client := helloService.NewHelloServiceClient(conn)
	r, err := client.SayHello(context.Background(), &helloService.Request{Data: []byte("send from client")})
	fmt.Printf("%s\n", string(r.Data))
}
