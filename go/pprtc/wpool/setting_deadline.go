package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	go performTask54(ctx)

	time.Sleep(3 * time.Second)
}

func performTask54(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task completed or deadline exceeded:", ctx.Err())
		return
	}
}

/*
	Setting timeouts and deadlines is crucial when working with context in Golang.
	It ensures that operations complete within a specified timeframe and prevents potential bottlenecks or indefinite waits.

	we create a context with a deadline of 2 seconds using context.WithDeadline(). The performTask goroutine waits for the context to be canceled or for the deadline to be exceeded.
	After 3 seconds, we let the program exit, triggering the deadline exceeded error */
