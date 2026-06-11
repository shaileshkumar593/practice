// example6_goroutine_lifetime/main.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// start periodic goroutine
	go periodic(ctx)

	// run for 1 second
	time.Sleep(1 * time.Second)

	// signal goroutine to stop
	fmt.Println("main: cancelling periodic worker")
	cancel()

	// allow goroutine to exit gracefully
	time.Sleep(200 * time.Millisecond)
}

func periodic(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop() // important to avoid timer leak

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-ctx.Done():
			fmt.Println("periodic: stopping:", ctx.Err())
			return
		}
	}
}

/*
	1. ticker.Stop()

Prevents OS timer objects from accumulating.

🟢 2. ctx.Done()

Notifies goroutine to exit.

🟢 3. Goroutine returns (not blocked on an infinite loop)

If you forget either:

Goroutine keeps running forever → goroutine leak

Timer stays active → timer leak

Service memory & CPU usage grows → microservice degradation

This is the correct industry pattern for microservices, Go workers, Kubernetes jobs.*/
