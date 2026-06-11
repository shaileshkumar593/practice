// example-2-manual-cancel/main.go

// Description: create WithCancel and cancel child goroutines when done.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "w1")
	go worker(ctx, "w2")

	time.Sleep(2 * time.Second)
	fmt.Println("main: canceling context")
	cancel() // signal workers to stop
	time.Sleep(1 * time.Second)
	fmt.Println("main: done")
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: received cancel: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s: working\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
