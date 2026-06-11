package main

import (
	"fmt"
	"time"
)

// Alternative: time.AfterFunc
func main() {
	// Create a new timer that will fire after 3 seconds
	// The program below creates a timer that waits for 3 seconds:
	timer := time.NewTimer(3 * time.Second)

	// Block and wait until the timer's channel receives a value
	<-timer.C
	fmt.Println("Timer fired after 3 seconds")
}
