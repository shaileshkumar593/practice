// example-4-deadline/main.go
// Description: stop work after a specified time T.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(700 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	if err := doStuff(ctx); err != nil {
		fmt.Println("stopped:", err)
	} else {
		fmt.Println("completed")
	}
}

func doStuff(ctx context.Context) error {
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(300 * time.Millisecond):
			fmt.Println("iteration", i)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

/*
	Behavior: stops at or before the deadline.
    Pitfall: deadlines can be hard to reason about across services; prefer timeouts unless absolute time needed.*/
