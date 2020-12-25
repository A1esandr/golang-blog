package main

import (
	"fmt"
	"sync"
)

type CountMu struct {
	mu      sync.Mutex
	counter int
}

// Example of syncing with channel and Mutexes

func main() {
	counter := &CountMu{}
	increment := func(c *CountMu, ch chan int) {
		c.mu.Lock()
		c.counter++
		ch <- 1
		c.mu.Unlock()
	}
	ch := make(chan int, 1)
	for i := 0; i < 20; i++ {
		go increment(counter, ch)
	}
	for i := 0; i < 20; i++ {
		<-ch
	}

	for i := 0; i < counter.counter; i++ {
		go func(i int) {
			fmt.Printf("Hello %d\n", i)
			ch <- 1
		}(i)
	}
	for i := 0; i < counter.counter; i++ {
		<-ch
	}
	fmt.Println(counter.counter)
}
