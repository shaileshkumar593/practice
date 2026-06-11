package main

import (
	"fmt"
	"time"
)

// sends and receives to an unbuffered channel are blocking.
func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	//time.Sleep(20 * time.Second)
	//time.Sleep(4 * time.Second)
	//fmt.Println(<-done)//
	/*Hello world goroutine
	main function */
	//fmt.Println("ok")
	//time.Sleep(10 * time.Second)
	fmt.Println("ohh")
	fmt.Println(<-done) // read from channel is always a blocking call until some write is done
	fmt.Println("wooo")
	//fmt.Println("bye")
}
func main() {
	done := make(chan bool)
	go hello(done)
	time.Sleep(10 * time.Second)
	done <- true

	fmt.Println("main function")
}
