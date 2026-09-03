package repository

import (
	"context"

	"grpc-order-service/internal/domain"
)

type OrderRepository interface {
	Create(
		ctx context.Context,
		order *domain.Order,
	) error

	Get(
		ctx context.Context,
		id string,
	) (*domain.Order, error)

	List(
		ctx context.Context,
		customerID string,
		limit int,
	) ([]*domain.Order, error)

	UpdateStatus(
		ctx context.Context,
		id string,
		status domain.OrderStatus,
	) error
}
