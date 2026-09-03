package order

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"grpc-order-service/internal/domain"
	"grpc-order-service/internal/repository"
)

type Service struct {
	repo repository.OrderRepository
}

func New(
	repo repository.OrderRepository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(
	ctx context.Context,
	customerID string,
	items []domain.OrderItem,
	currency string,
) (*domain.Order, error) {

	customerID = strings.TrimSpace(customerID)
	currency = strings.TrimSpace(currency)

	if customerID == "" {
		return nil, fmt.Errorf(
			"customer ID is required",
		)
	}

	if len(items) == 0 {
		return nil, fmt.Errorf(
			"at least one item is required",
		)
	}

	if currency == "" {
		return nil, fmt.Errorf(
			"currency is required",
		)
	}

	var total float64

	for _, item := range items {

		if item.Quantity <= 0 {
			return nil, fmt.Errorf(
				"quantity must be positive",
			)
		}

		if item.Price < 0 {
			return nil, fmt.Errorf(
				"price cannot be negative",
			)
		}

		total +=
			float64(item.Quantity) *
				item.Price
	}

	now := time.Now().UTC()

	order := &domain.Order{
		ID:          uuid.NewString(),
		CustomerID:  customerID,
		Items:       items,
		TotalAmount: total,
		Currency:    currency,
		Status:      domain.OrderPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(
		ctx,
		order,
	); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Service) Get(
	ctx context.Context,
	id string,
) (*domain.Order, error) {

	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf(
			"order ID is required",
		)
	}

	return s.repo.Get(ctx, id)
}

func (s *Service) Cancel(
	ctx context.Context,
	id string,
) (*domain.Order, error) {

	order, err := s.repo.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if order.Status == domain.OrderCancelled {
		return order, nil
	}

	if order.Status == domain.OrderConfirmed {
		return nil, fmt.Errorf(
			"confirmed order cannot be cancelled",
		)
	}

	err = s.repo.UpdateStatus(
		ctx,
		id,
		domain.OrderCancelled,
	)

	if err != nil {
		return nil, err
	}

	order.Status = domain.OrderCancelled

	return order, nil
}
