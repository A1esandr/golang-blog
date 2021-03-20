package main

import (
	"fmt"
	"strconv"
)

func main() {

	channelsWithoutLabels()
	//channels()
	//loop()

}

func loop() {
	a := 0
label:
	for i := 0; i < 10; i++ {
		if a > 10 {
			fmt.Println("After a")
			break
		}
		if i%2 == 0 {
			a++
			fmt.Println(strconv.Itoa(i))
			break label
		}
	}

	fmt.Println("After")

}

func channels() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
	label:
		for {
			select {
			case a := <-ch1:
				fmt.Println("ch1")
				if a == 2 {
					break label
				}
				fmt.Println(strconv.Itoa(a))
			case b := <-ch2:
				fmt.Println("ch2")
				fmt.Println(strconv.Itoa(b))
			}
			fmt.Println("Out of select")
		}
		fmt.Println("Out of for")
	}()
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Printf("Send %d\n", i)
		ch1 <- i
	}

	fmt.Println("End")
}

func channelsWithoutLabels() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 1; {
			select {
			case a := <-ch1:
				fmt.Println("ch1")
				if a == 2 {
					i++
					break
				}
				fmt.Println(strconv.Itoa(a))
			case b := <-ch2:
				fmt.Println("ch2")
				fmt.Println(strconv.Itoa(b))
			}
			fmt.Println("Out of select")
		}
		fmt.Println("Out of for")
	}()
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Printf("Send %d\n", i)
		ch1 <- i
	}

	fmt.Println("End")
}
