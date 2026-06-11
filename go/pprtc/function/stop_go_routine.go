package main

import (
	"fmt"
	"time"
)

func worker(stop chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("Worker stopping...")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop := make(chan struct{})

	go worker(stop)

	time.Sleep(3 * time.Second)

	close(stop) // signal stop

	time.Sleep(1 * time.Second)
}

/*

	Why use struct{}?
Zero memory

Used only as signal

Very efficient

*/
