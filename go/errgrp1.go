package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func fetchUser(ctx context.Context) (string, error) {

	select {
	case <-time.After(2 * time.Second):
		return "Shailesh", nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func fetchBooking(ctx context.Context) (string, error) {

	select {
	case <-time.After(3 * time.Second):
		return "Booking-123", nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func fetchPayment(ctx context.Context) (string, error) {

	select {
	case <-time.After(1 * time.Second):
		return "Paid", nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	var (
		user    string
		booking string
		payment string
	)

	g.Go(func() error {

		u, err := fetchUser(ctx)
		if err != nil {
			return err
		}

		user = u
		return nil
	})

	g.Go(func() error {

		b, err := fetchBooking(ctx)
		if err != nil {
			return err
		}

		booking = b
		return nil
	})

	g.Go(func() error {

		p, err := fetchPayment(ctx)
		if err != nil {
			return err
		}

		payment = p
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User:", user)
	fmt.Println("Booking:", booking)
	fmt.Println("Payment:", payment)
}
