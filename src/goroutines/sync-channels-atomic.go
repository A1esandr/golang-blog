package main

import (
	"fmt"
	"sync/atomic"
)

// Example of syncing with channel and atomic

func main() {
	var counter int64
	increment := func(c *int64, ch chan int) {
		atomic.AddInt64(c, 1)
		ch <- 1
	}
	ch := make(chan int, 1)
	for i := 0; i < 20; i++ {
		go increment(&counter, ch)
	}
	for i := 0; i < 20; i++ {
		<-ch
	}

	var i int64
	for i = 0; i < counter; i++ {
		go func(i int64) {
			fmt.Printf("Hello %d\n", i)
			ch <- 1
		}(i)
	}
	for i = 0; i < counter; i++ {
		<-ch
	}
	fmt.Println(counter)
}
