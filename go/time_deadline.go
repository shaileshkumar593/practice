package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Simulate an expensive operation
func ProcessOrder(ctx context.Context) error {

	// Check whether the context has a deadline
	deadline, ok := ctx.Deadline()

	if ok {
		fmt.Println("Deadline:", deadline.Format("15:04:05"))

		remaining := time.Until(deadline)

		fmt.Println("Time Remaining:", remaining)

		// If less than 1 second remains, don't even start.
		if remaining < time.Second {
			return errors.New("not enough time remaining")
		}
	}

	fmt.Println("Starting expensive operation...")

	// Simulate work that takes 3 seconds
	select {

	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed successfully.")
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {

	// Context expires after only 500 milliseconds
	ctx, cancel := context.WithTimeout(
		context.Background(),
		500*time.Millisecond,
	)
	defer cancel()

	err := ProcessOrder(ctx)

	if err != nil {
		fmt.Println("Error:", err)
	}
}