package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return processUser(ctx)
	})

	g.Go(func() error {
		return processOrder(ctx)
	})

	g.Go(func() error {
		return processPayment(ctx)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Operation failed:", err)
	}
}

func processUser(ctx context.Context) error {
	return nil
}

func processOrder(ctx context.Context) error {
	return fmt.Errorf("order failed")
}

func processPayment(ctx context.Context) error {
	return nil
}
