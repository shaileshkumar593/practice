package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	go performTask(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Task timed out")
	}
}

func performTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
	}
}

/*

	In this example, the performTask function simulates a long-running task that takes 5 seconds to complete.
	However, since the context has a timeout of only 2 seconds, the operation is terminated prematurely, resulting in a timeout.*/
