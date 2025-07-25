package order

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/order"
	entity "post-tech-challenge-10soat/internal/entities"
)

type CreateOrderUseCase interface {
	Execute(ctx context.Context, createOrder dto.CreateOrderDTO) (entity.Order, error)
}
