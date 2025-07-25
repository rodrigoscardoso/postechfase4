package interfaces

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type OrderGateway interface {
	CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error)
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context, limit uint64) ([]entity.Order, error)
	GetOrderById(ctx context.Context, id string) (entity.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) (entity.Order, error)
}
