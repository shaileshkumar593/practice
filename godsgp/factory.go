package main

import "fmt"

type Payment interface {
	Pay(amount float64)
}

type Stripe struct{}

func (s Stripe) Pay(amount float64) {
	fmt.Println("Stripe payment:", amount)
}

type PayPal struct{}

func (p PayPal) Pay(amount float64) {
	fmt.Println("PayPal payment:", amount)
}

// Factory hides object creation from the caller.
func NewPayment(provider string) Payment {
	switch provider {
	case "stripe":
		return Stripe{}

	case "paypal":
		return PayPal{}

	default:
		return nil
	}
}

func main() {
	payment := NewPayment("stripe")

	payment.Pay(100)
}
