package main

import "fmt"

func hello(p chan int) {
	fmt.Println("ok in goroutine")
}

func main() {
	ch := make(chan int)
	go hello(ch)
	ch <- 5
	// if write on channel then there must be
	//some read happen otherwise channel is blocked and cause deadlock
}
