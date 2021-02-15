package main

import "fmt"

type semafor chan struct{}

func New(n int) semafor {
	return make(semafor, n)
}

func (s semafor) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s semafor) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	sem := New(n)

	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}

	sem.Dec(n)

}
