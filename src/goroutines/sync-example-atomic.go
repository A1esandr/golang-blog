package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Example of syncing with WaitGroup and atomic package

func main() {

	var counter int64
	increment := func(c *int64) {
		atomic.AddInt64(c, 1)
	}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			increment(&counter)
			wg.Done()
		}()

	}
	wg.Wait()

	var i int64
	for i = 0; i < counter; i++ {
		wg.Add(1)
		go func(i int64) {
			fmt.Printf("Hello %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(counter)
}
