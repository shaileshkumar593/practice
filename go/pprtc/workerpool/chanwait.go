package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	time.Sleep(10 * time.Second)
	done <- true
}
func main() {
	done := make(chan bool)
	go hello(done)
	fmt.Println(<-done) // here progm control is wating for data
	fmt.Println("main function")
}
