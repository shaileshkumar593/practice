package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64)
}

type LegacyPaymentSystem struct{}

func (l LegacyPaymentSystem) MakePayment(amount float64) {
	fmt.Println("Legacy payment:", amount)
}

type PaymentAdapter struct {
	legacy LegacyPaymentSystem
}

func (a PaymentAdapter) Pay(amount float64) {
	a.legacy.MakePayment(amount)
}

func main() {
	legacy := LegacyPaymentSystem{}

	adapter := PaymentAdapter{
		legacy: legacy,
	}

	var processor PaymentProcessor = adapter

	processor.Pay(500)
}
