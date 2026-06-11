// example-10-long-task/main.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 750*time.Millisecond)
	defer cancel()
	if err := processAll(ctx); err != nil {
		fmt.Println("processAll stopped:", err)
	}
}

func processAll(ctx context.Context) error {
	items := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		items = append(items, i)
	}
	for _, it := range items {
		// simulate small unit of work
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		_ = it // process
	}
	return nil
}

// Description: break long tasks into chunks and check ctx.Done() between chunks.
// Pitfall: check context frequently enough (per loop) but not excessively.
