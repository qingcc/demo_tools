package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("aaa")
	for {
		var buffer [512]byte
		n, err := os.Stdin.Read(buffer[:])
		if err != nil {
			fmt.Println("read error:", err)
			return
		}
		aa := string(buffer[:n])
		fmt.Printf("input:%s", aa)
	}
}
