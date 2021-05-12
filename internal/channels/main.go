package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		read(ch)
		wg.Done()
	}()
	go func() {
		write(ch)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finish")
}

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("Wrote", i)
	}
}

func read(ch chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println("Read", <-ch)
	}
}
