package main

import "fmt"

// Receiving from channel will wait until channel have closed
func main() {
	w := make(chan struct{})

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		close(w)
	}()
	<-w
	fmt.Println("Finish")
}
