package order

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type ListOrdersUseCase interface {
	Execute(ctx context.Context, limit uint64) ([]entity.Order, error)
}
