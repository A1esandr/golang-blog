package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var count int64

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddInt64(&count, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", count)
}
