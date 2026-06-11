package main

import (
	"fmt"
)

func producer1(ch chan int) {

	defer close(ch)

	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func main() {

	ch := make(chan int)

	go producer1(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
