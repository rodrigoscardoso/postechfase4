package product

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type ListProductsUseCase interface {
	Execute(ctx context.Context, categoryId string) ([]entity.Product, error)
}
