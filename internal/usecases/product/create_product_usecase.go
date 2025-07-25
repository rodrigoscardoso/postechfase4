package product

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/product"
	entity "post-tech-challenge-10soat/internal/entities"
)

type CreateProductUseCase interface {
	Execute(ctx context.Context, product dto.CreateProductDTO) (entity.Product, error)
}
