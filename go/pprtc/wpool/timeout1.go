package main

import (
	"fmt"
	"time"
)

func longRunningOperation(resultChan chan string) {
	time.Sleep(3 * time.Second) // Simulate a long operation
	resultChan <- "Operation Complete"
}

func main() {
	resultChan := make(chan string, 1)
	go longRunningOperation(resultChan)

	select {
	case result := <-resultChan:
		fmt.Println(result)
	case <-time.After(2 * time.Second): // Timeout after 2 seconds
		fmt.Println("Operation timed out!")
	}
}

/*
This Go program demonstrates how to handle timeouts when waiting for a long-running operation to complete.
The `longRunningOperation` function simulates a task that takes 3 seconds to finish. In the `main` function,
we start this operation in a separate goroutine and use a `select` statement to wait for either the result of
the operation or a timeout of 2 seconds.For one-off delays, time.After works well. However, in a for loop,
repeatedly calling time.After can cause a memory leak if the timer is created but the case is never selected
(e.g., if the main operation finishes first). The underlying time.Timer is not garbage collected until it fires.
To avoid this in loops or long-running applications, use time.NewTimer and explicitly call its Stop method when
the timer is no longer needed: */
