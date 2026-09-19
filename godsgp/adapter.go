package main

import "fmt"

// =====================================================
// 1. Target Interface
// =====================================================

// This is what our application expects.
type PaymentGateway interface {
	Pay(amount float64) error
}

// =====================================================
// 2. Third-party Stripe SDK
// =====================================================

// Imagine this comes from Stripe's SDK.
// We cannot modify this code.
type StripeSDK struct {
	APIKey string
}

func (s *StripeSDK) CreatePayment(amount float64) {
	fmt.Printf("Stripe SDK: Creating payment of ₹%.2f\n", amount)
}

// =====================================================
// 3. Third-party Razorpay SDK
// =====================================================

// Imagine this comes from Razorpay's SDK.
// We cannot modify this code.
type RazorpaySDK struct {
	Key string
}

func (r *RazorpaySDK) CreateOrder(amount float64) {
	fmt.Printf("Razorpay SDK: Creating order of ₹%.2f\n", amount)
}

// =====================================================
// 4. Stripe Adapter
// =====================================================

type StripeAdapter struct {
	stripe *StripeSDK
}

// Adapter implements PaymentGateway.
func (s *StripeAdapter) Pay(amount float64) error {

	fmt.Println("StripeAdapter: Converting application request")

	s.stripe.CreatePayment(amount)

	return nil
}

// =====================================================
// 5. Razorpay Adapter
// =====================================================

type RazorpayAdapter struct {
	razorpay *RazorpaySDK
}

// Adapter implements PaymentGateway.
func (r *RazorpayAdapter) Pay(amount float64) error {

	fmt.Println("RazorpayAdapter: Converting application request")

	r.razorpay.CreateOrder(amount)

	return nil
}

// =====================================================
// 6. Payment Service
// =====================================================

type PaymentService struct {
	gateway PaymentGateway
}

func NewPaymentService(gateway PaymentGateway) *PaymentService {
	return &PaymentService{
		gateway: gateway,
	}
}

func (p *PaymentService) ProcessPayment(amount float64) error {

	fmt.Printf("PaymentService: Processing ₹%.2f\n", amount)

	return p.gateway.Pay(amount)
}

// =====================================================
// 7. Main
// =====================================================

func main() {

	// -------------------------------
	// Stripe
	// -------------------------------

	stripeSDK := &StripeSDK{
		APIKey: "stripe-secret",
	}

	stripeAdapter := &StripeAdapter{
		stripe: stripeSDK,
	}

	stripeService := NewPaymentService(stripeAdapter)

	err := stripeService.ProcessPayment(1000)

	if err != nil {
		fmt.Println("Payment failed:", err)
	}

	fmt.Println()

	// -------------------------------
	// Razorpay
	// -------------------------------

	razorpaySDK := &RazorpaySDK{
		Key: "razorpay-secret",
	}

	razorpayAdapter := &RazorpayAdapter{
		razorpay: razorpaySDK,
	}

	razorpayService := NewPaymentService(razorpayAdapter)

	err = razorpayService.ProcessPayment(2000)

	if err != nil {
		fmt.Println("Payment failed:", err)
	}
}

/*

	Scenario
Our application expects every payment provider to implement:

Pay(amount float64) error
But third-party SDKs have different methods:

Stripe → CreatePayment()

Razorpay → CreateOrder()
*/
