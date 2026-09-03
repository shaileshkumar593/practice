package domain

import "time"

type OrderStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderConfirmed OrderStatus = "CONFIRMED"
	OrderCancelled OrderStatus = "CANCELLED"
)

type OrderItem struct {
	ProductID string
	Quantity  int32
	Price     float64
}

type Order struct {
	ID          string
	CustomerID  string
	Items       []OrderItem
	TotalAmount float64
	Currency    string
	Status      OrderStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
