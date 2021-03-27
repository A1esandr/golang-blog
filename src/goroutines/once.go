package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(func() {
				fmt.Printf("Only once %v \n", i)
			})
			fmt.Printf("%v \n", i)
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
