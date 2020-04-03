package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func gen(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			select {
			case out <- num:
				log.Println("write:", num)
			case <-done:
				log.Printf("gen:done")
				return
			}
		}
		close(out)
	}()
	return out
}

func sq(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			select {
			case out <- n * n:
				log.Println("sq:", n)
				//case <-done:
				//	log.Printf("sq:done")
				return
			}
		}
		close(out)
	}()
	return out
}

func merge(done chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	output := func(in <-chan int) {
		for n := range in {
			select {
			case out <- n:
				log.Println("out:", n)
				//case <-done:
				//	log.Printf("merge:done")
				return
			}
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	done := make(chan struct{})
	c := gen(done, 1, 2, 3)
	out1 := sq(done, c)
	out2 := sq(done, c)
	out := merge(done, out1, out2)
	//for n := range out {
	//	fmt.Println(n)
	//}
	fmt.Println(<-out)
	done <- struct{}{}
	//time.Sleep(time.Second*3)
	log.Println("over")
	for {
		time.Sleep(time.Second)
	}
}
