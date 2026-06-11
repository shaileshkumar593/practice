package main

import (
	"context"
	"errors"
	"fmt" // Assuming we might use this context for real calls
	"time"

	// Import the errgroup package
	"golang.org/x/sync/errgroup"
)

// Product model remains the same
type Product struct {
	ID           int
	Availability string
	Price        float64
	// Use pointers or sync.Mutex for concurrent writes if needed,
	// but simple fields are okay for this demo if reads happen after Wait().
}

// Simulate fetching availability
func fetchAvailability(ctx context.Context, product *Product) error {
	// Simulate network call that respects context
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Fetching availability for product 1...")
		product.Availability = "In Stock"
		fmt.Println("Availability fetched.")
		return nil
	case <-ctx.Done():
		fmt.Println("Availability fetch cancelled:", ctx.Err())
		return ctx.Err() // Propagate context error
	}
}

// Simulate fetching price - this one will fail
func fetchPrice(ctx context.Context, product *Product) error {
	// Simulate network call that respects context but returns an error
	select {
	case <-time.After(50 * time.Millisecond): // Faster, but will error
		fmt.Println("Fetching price for product 1...")
		fmt.Println("Price fetch failed!")
		return errors.New("simulated network error fetching price")
	case <-ctx.Done():
		fmt.Println("Price fetch cancelled:", ctx.Err())
		return ctx.Err()
	}
}

func main() {
	product := &Product{ID: 1}

	// Create an errgroup.Group together with a cancellable context.
	// If any function run by eg.Go returns an error, eg.Wait() will return that error,
	// AND the context 'ctx' will be cancelled.
	eg, ctx := errgroup.WithContext(context.Background())

	// Goroutine to fetch availability
	eg.Go(func() error {
		// Pass the derived context 'ctx' down.
		// The function itself handles context cancellation.
		return fetchAvailability(ctx, product)
		// No wg.Add or wg.Done needed! errgroup handles the counting.
	})

	// Goroutine to fetch price
	eg.Go(func() error {
		// Pass the derived context 'ctx' down.
		return fetchPrice(ctx, product)
	})

	// Wait for both goroutines to complete.
	fmt.Println("Router is waiting for subgraph responses...")
	// eg.Wait() blocks until all goroutines launched by eg.Go have returned.
	// It returns the first non-nil error encountered, or nil if all succeed.
	// Importantly, if one eg.Go function returns an error, the 'ctx' derived
	// from errgroup.WithContext is cancelled, signaling other running goroutines
	// (that respect the context) to stop early.
	err := eg.Wait()

	// Check the error returned by Wait()
	if err != nil {
		fmt.Println("Error occurred during fetch:", err)
		// Because fetchPrice failed, ctx was cancelled. Depending on timing,
		// fetchAvailability might have printed "Availability fetch cancelled".
		// The product data might be incomplete.
	} else {
		// Only print the full product if everything succeeded.
		fmt.Printf("Successfully fetched data for Product ID %d: %+v\n", product.ID, *product)
	}
}
