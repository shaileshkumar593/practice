package service

import (
	"context"
	"fmt"
	"time"

	proto "github.com/example/order-service/proto"
	"github.com/example/order-service/internal/store"
)

// OrderService implements proto.OrderServiceServer
type OrderService struct {
	proto.UnimplementedOrderServiceServer
	store *store.InMemoryStore
}

func NewOrderService(s *store.InMemoryStore) *OrderService {
	return &OrderService{store: s}
}

func (s *OrderService) GetOrder(ctx context.Context, req *proto.GetOrderRequest) (*proto.GetOrderResponse, error) {
	if req == nil || req.Id == "" {
		return nil, fmt.Errorf("id required")
	}
	ord, err := s.store.Get(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetOrderResponse{Order: toProto(ord)}, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("invalid request")
	}
	ord, err := s.store.Create(req.Items, req.Price)
	if err != nil {
		return nil, err
	}
	return &proto.CreateOrderResponse{Order: toProto(ord)}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *proto.UpdateOrderRequest) (*proto.UpdateOrderResponse, error) {
	if req == nil || req.Id == "" {
		return nil, fmt.Errorf("id required")
	}
	ord, err := s.store.Update(req.Id, req.Items, req.Price, req.Status)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateOrderResponse{Order: toProto(ord)}, nil
}

func toProto(o store.Order) *proto.Order {
	return &proto.Order{
		Id:        o.ID,
		Items:     o.Items,
		Price:     o.Price,
		Status:    o.Status,
		CreatedAt: o.CreatedAt.Format(time.RFC3339),
		UpdatedAt: o.UpdatedAt.Format(time.RFC3339),
	}
}
