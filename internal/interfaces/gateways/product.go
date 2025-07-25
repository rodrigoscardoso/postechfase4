package interfaces

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type ProductGateway interface {
	ListProducts(ctx context.Context, categoryId string) ([]entity.Product, error)
	GetProductById(ctx context.Context, id string) (entity.Product, error)
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	UpdateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}
