package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type User struct {
	ID   int
	Name string
}

type Booking struct {
	ID     int
	Amount float64
}

type Payment struct {
	ID     int
	Status string
}

type Supplier struct {
	ID   int
	Name string
}

type BookingResponse struct {
	User     User
	Booking  Booking
	Payment  Payment
	Supplier Supplier
}

func fetchUser(ctx context.Context, userID int) (User, error) {

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("User fetched")
		return User{
			ID:   userID,
			Name: "Shailesh",
		}, nil

	case <-ctx.Done():
		fmt.Println("User cancelled")
		return User{}, ctx.Err()
	}
}

func fetchBooking(ctx context.Context, bookingID int) (Booking, error) {

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Booking fetched")
		return Booking{
			ID:     bookingID,
			Amount: 1000,
		}, nil

	case <-ctx.Done():
		fmt.Println("Booking cancelled")
		return Booking{}, ctx.Err()
	}
}

func fetchPayment(ctx context.Context, bookingID int) (Payment, error) {

	select {
	case <-time.After(1 * time.Second):

		// Simulate failure
		return Payment{},
			errors.New("payment service unavailable")

	case <-ctx.Done():
		fmt.Println("Payment cancelled")
		return Payment{}, ctx.Err()
	}
}

func fetchSupplier(ctx context.Context, bookingID int) (Supplier, error) {

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Supplier fetched")

		return Supplier{
			ID:   1,
			Name: "ABC Travels",
		}, nil

	case <-ctx.Done():
		fmt.Println("Supplier cancelled")
		return Supplier{}, ctx.Err()
	}
}

func GetBookingDetails(
	ctx context.Context,
	bookingID int,
) (*BookingResponse, error) {

	g, ctx := errgroup.WithContext(ctx)

	var (
		user     User
		booking  Booking
		payment  Payment
		supplier Supplier
	)

	g.Go(func() error {

		u, err := fetchUser(ctx, 1)
		if err != nil {
			return fmt.Errorf(
				"fetch user: %w",
				err,
			)
		}

		user = u
		return nil
	})

	g.Go(func() error {

		b, err := fetchBooking(ctx, bookingID)
		if err != nil {
			return fmt.Errorf(
				"fetch booking: %w",
				err,
			)
		}

		booking = b
		return nil
	})

	g.Go(func() error {

		p, err := fetchPayment(ctx, bookingID)
		if err != nil {
			return fmt.Errorf(
				"fetch payment: %w",
				err,
			)
		}

		payment = p
		return nil
	})

	g.Go(func() error {

		s, err := fetchSupplier(ctx, bookingID)
		if err != nil {
			return fmt.Errorf(
				"fetch supplier: %w",
				err,
			)
		}

		supplier = s
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &BookingResponse{
		User:     user,
		Booking:  booking,
		Payment:  payment,
		Supplier: supplier,
	}, nil
}

func main() {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	resp, err := GetBookingDetails(ctx, 123)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("%+v\n", resp)
}
