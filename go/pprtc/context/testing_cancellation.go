// example-11-testing/main.go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // immediate cancel
	if err := operation(ctx); err != nil {
		fmt.Println("operation returned:", err)
	}
}

func operation(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// unreachable here since ctx already canceled
	}
	return nil
}

//  Description: tests can create canceled contexts to ensure code responds to cancellation.
//  Pitfall: in unit tests, avoid flakiness by using deterministic sleeps or channels.
