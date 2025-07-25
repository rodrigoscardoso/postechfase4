package order

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type UpdateOrderStatusUseCase interface {
	Execute(ctx context.Context, id string, status string) (entity.Order, error)
}
