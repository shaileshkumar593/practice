package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := time.Tick(500 * time.Millisecond)

	// func Tick(d Duration) <-chan Time
	/*

		So you cannot:

			Stop it

			Release resources

			Control lifecycle
	*/
	request := []int{1, 2, 3, 4, 5}

	for _, req := range request {
		<-limiter

		fmt.Println("Processing request ", req, time.Now())
	}
}
