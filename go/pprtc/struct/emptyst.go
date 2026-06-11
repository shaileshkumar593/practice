package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	go func() {
		// Simulate work
		fmt.Println("Working...")
		time.Sleep(3 * time.Second)
		// Send quit signal
		close(quit)
	}()

	// Block and wait for quit signal to close
	<-quit
	fmt.Println("Quit signal received, exiting...")
}
