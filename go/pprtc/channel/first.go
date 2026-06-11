package main

import "fmt"

func MyFunc(ch chan int) {
	ch <- (45 + <-ch)
}

func main() {
	fmt.Println("hello First channel program ")
	ch := make(chan int)

	go MyFunc(ch)
	ch <- 78

	fmt.Println("Value from channel is ", <-ch)
}
