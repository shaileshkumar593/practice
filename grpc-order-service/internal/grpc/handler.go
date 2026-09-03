package grpcserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	orderv1 "grpc-order-service/api/gen/order/v1"
	apporder "grpc-order-service/internal/application/order"
	"grpc-order-service/internal/domain"
)

type OrderHandler struct {
	orderv1.UnimplementedOrderServiceServer

	service *apporder.Service
}

func NewOrderHandler(
	service *apporder.Service,
) *OrderHandler {

	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(
	ctx context.Context,
	req *orderv1.CreateOrderRequest,
) (*orderv1.CreateOrderResponse, error) {

	if req == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"request is required",
		)
	}

	items := make(
		[]domain.OrderItem,
		0,
		len(req.Items),
	)

	for _, item := range req.Items {

		items = append(
			items,
			domain.OrderItem{
				ProductID: item.ProductId,
				Quantity:  item.Quantity,
				Price:     item.Price,
			},
		)
	}

	order, err := h.service.Create(
		ctx,
		req.CustomerId,
		items,
		req.Currency,
	)

	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	return &orderv1.CreateOrderResponse{
		Order: toProto(order),
	}, nil
}

func (h *OrderHandler) GetOrder(
	ctx context.Context,
	req *orderv1.GetOrderRequest,
) (*orderv1.GetOrderResponse, error) {

	if req == nil || req.OrderId == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"order_id is required",
		)
	}

	order, err := h.service.Get(
		ctx,
		req.OrderId,
	)

	if err != nil {
		return nil, status.Error(
			codes.NotFound,
			err.Error(),
		)
	}

	return &orderv1.GetOrderResponse{
		Order: toProto(order),
	}, nil
}

func toProto(
	order *domain.Order,
) *orderv1.Order {

	if order == nil {
		return nil
	}

	items := make(
		[]*orderv1.OrderItem,
		0,
		len(order.Items),
	)

	for _, item := range order.Items {

		items = append(
			items,
			&orderv1.OrderItem{
				ProductId: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
			},
		)
	}

	return &orderv1.Order{
		Id:          order.ID,
		CustomerId:  order.CustomerID,
		Items:       items,
		TotalAmount: order.TotalAmount,
		Currency:    order.Currency,
		Status:      toProtoStatus(order.Status),
	}
}
