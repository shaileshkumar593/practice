// example-6-goroutine-lifetime/main.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go periodic(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(200 * time.Millisecond)
}

func periodic(ctx context.Context) {
	t := time.NewTicker(200 * time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("tick")
		case <-ctx.Done():
			fmt.Println("periodic: stopping:", ctx.Err())
			return
		}
	}
}

// Pitfall: make sure all goroutines that could outlive the request observe ctx.Done().
