package main

import (
	"fmt"
	"sync"
)

type Count struct {
	mu      sync.Mutex
	counter int
}

// Example of syncing with WaitGroup and Mutexes

func main() {
	counter := &Count{}
	increment := func(c *Count) {
		c.mu.Lock()
		c.counter++
		c.mu.Unlock()
	}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			increment(counter)
			wg.Done()
		}()

	}
	wg.Wait()

	for i := 0; i < counter.counter; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Hello %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(counter.counter)
}
