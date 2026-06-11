package main

import (
	"fmt"
	"time"
)

func main() {
	// AfterFunc calls a function when the timer expires, no channel read needed
	// i.e., it executes the callback function after the specified duration
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("Callback function executed")
	})

	// Give the callback time to run, then stop the main program
	time.Sleep(3 * time.Second)
	timer.Stop() // Optional: stop the timer (though it already ran)
}
