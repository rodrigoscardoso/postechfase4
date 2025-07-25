package interfaces

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/order"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order dto.CreateOrderDTO) (dto.OrderDTO, error)
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context, limit uint64) ([]dto.OrderDTO, error)
	GetOrderById(ctx context.Context, id string) (dto.OrderDTO, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) (dto.OrderDTO, error)
}
