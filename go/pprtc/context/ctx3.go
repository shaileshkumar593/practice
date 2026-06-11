// example-3-timeout/main.go
// Description: bound operations to a max duration.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// set a short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	err := doWork(ctx)
	if err != nil {
		fmt.Println("main:", err)
		return
	}
	fmt.Println("main: work completed")
}

func doWork(ctx context.Context) error {
	// simulate a slow operation that takes 1s
	select {
	case <-time.After(1 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err() // context deadline exceeded
	}
}

/*
	Behavior: returns context deadline exceeded.
    Pitfall: always defer cancel() to release timer resources.
*/
