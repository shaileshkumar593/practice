package main

import (
	"fmt"
)

func producer(ch chan int) {

	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch) // required because range on ch required channel to be close
}

func consumer(ch chan int) {

	for val := range ch {
		fmt.Println("Consumed:", val)
	}
}

func main() {

	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}
