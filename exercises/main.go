package main

import (
	"exercises/concurrency"
	"fmt"
	"sync"
	"time"
)

// Buffered and Unbuffered Channels
func send(c chan int, wg *sync.WaitGroup) {
	// defer wg.Done()
	c <- 1
	c <- 2
}

func channel() {
	c := make(chan int)
	var wg sync.WaitGroup
	// wg.Add(1)
	go send(c, &wg)
	v := <-c
	// wg.Wait()
	time.Sleep(time.Second * 3)
	fmt.Println(v)
}

func main() {
	// channel()

	concurrency.Concurrency(5)
}
