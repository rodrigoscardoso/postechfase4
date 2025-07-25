package interfaces

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type OrderProductGateway interface {
	CreateOrderProduct(ctx context.Context, orderProduct entity.OrderProduct) (entity.OrderProduct, error)
}
