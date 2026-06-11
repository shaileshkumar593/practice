package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Stopping worker")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan struct{})

	go worker(done)

	time.Sleep(2 * time.Second)

	close(done) // signal stop
	time.Sleep(1 * time.Second)
}

/*

Channel Signaling (Goroutine Stop Signal)
Why not chan bool?

Because:

No data needed

Only signal

Zero memory

Cleaner semantic meaning

*/
