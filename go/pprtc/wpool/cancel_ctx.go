package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go performTaskCancel(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}

func performTaskCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			// Perform task operation
			fmt.Println("Performing task...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
	Cancellation is an essential aspect of context management. It allows you to gracefully terminate operations and propagate
	cancellation signals to related goroutines. By canceling a context, you can avoid resource leaks and ensure the timely
	termination of concurrent operations.

	In this example, we create a context using context.WithCancel() and defer the cancellation function.
	The performTask goroutine continuously performs a task until the context is canceled. After 2 seconds,
	we call the cancel function to initiate the cancellation process. As a result, the goroutine detects the
	cancellation signal and terminates the task.*/
