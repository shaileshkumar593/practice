package main

import "fmt"

// =====================================================
// 1. Strategy Interface
// =====================================================

type PaymentStrategy interface {
	Pay(amount float64) error
}

// =====================================================
// 2. Card Strategy
// =====================================================

type CardPayment struct {
	CardNumber string
}

func (c *CardPayment) Pay(amount float64) error {

	fmt.Printf(
		"CardPayment: Charging ₹%.2f using card %s\n",
		amount,
		c.CardNumber,
	)

	return nil
}

// =====================================================
// 3. UPI Strategy
// =====================================================

type UPIPayment struct {
	UPIID string
}

func (u *UPIPayment) Pay(amount float64) error {

	fmt.Printf(
		"UPIPayment: Charging ₹%.2f using UPI %s\n",
		amount,
		u.UPIID,
	)

	return nil
}

// =====================================================
// 4. Wallet Strategy
// =====================================================

type WalletPayment struct {
	WalletID string
}

func (w *WalletPayment) Pay(amount float64) error {

	fmt.Printf(
		"WalletPayment: Charging ₹%.2f from wallet %s\n",
		amount,
		w.WalletID,
	)

	return nil
}

// =====================================================
// 5. Context
// =====================================================

type PaymentService struct {
	strategy PaymentStrategy
}

// Constructor

func NewPaymentService(strategy PaymentStrategy) *PaymentService {
	return &PaymentService{
		strategy: strategy,
	}
}

// Change strategy at runtime

func (p *PaymentService) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// Execute selected strategy

func (p *PaymentService) ProcessPayment(amount float64) error {

	if p.strategy == nil {
		return fmt.Errorf("payment strategy is not configured")
	}

	return p.strategy.Pay(amount)
}

// =====================================================
// 6. Main
// =====================================================

func main() {

	// -------------------------------
	// Card Payment
	// -------------------------------

	card := &CardPayment{
		CardNumber: "XXXX-XXXX-XXXX-1234",
	}

	service := NewPaymentService(card)

	err := service.ProcessPayment(1000)

	if err != nil {
		fmt.Println("Payment failed:", err)
	}

	fmt.Println()

	// -------------------------------
	// Change strategy to UPI
	// -------------------------------

	upi := &UPIPayment{
		UPIID: "user@upi",
	}

	service.SetStrategy(upi)

	err = service.ProcessPayment(2000)

	if err != nil {
		fmt.Println("Payment failed:", err)
	}

	fmt.Println()

	// -------------------------------
	// Change strategy to Wallet
	// -------------------------------

	wallet := &WalletPayment{
		WalletID: "wallet-123",
	}

	service.SetStrategy(wallet)

	err = service.ProcessPayment(3000)

	if err != nil {
		fmt.Println("Payment failed:", err)
	}
}

/*

	Our application supports multiple payment algorithms:

Card
UPI
Wallet
We want to select the payment method at runtime.
*/
