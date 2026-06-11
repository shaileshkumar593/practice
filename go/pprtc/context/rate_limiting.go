// example-9-rate-limit/main.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	limiter := time.Tick(200 * time.Millisecond) // simple ticker limiter
	ctx, cancel := context.WithTimeout(context.Background(), 900*time.Millisecond)
	defer cancel()
	for i := 0; i < 10; i++ {
		select {
		case <-limiter:
			fmt.Println("do work", i)
		case <-ctx.Done():
			fmt.Println("stopped by context:", ctx.Err())
			return
		}
	}
}

// Description: combine rate limiter and context to stop waiting when canceled.
// Pitfall: prefer robust token bucket libraries for production.
