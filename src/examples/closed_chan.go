package main

import "fmt"

// Receiving from channel will wait until channel have closed
func main() {
	not_buffered()
	buffered()
}

func not_buffered() {
	fmt.Println("not_buffered")
	w := make(chan struct{})

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		close(w)
	}()
	<-w
	fmt.Println("Finish not_buffered")
}

func buffered() {
	fmt.Println("buffered")
	w := make(chan struct{}, 10)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		close(w)
	}()
	<-w
	fmt.Println("Finish buffered")
}
