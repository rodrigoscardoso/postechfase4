package interfaces

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/product"
)

type ProductRepository interface {
	ListProducts(ctx context.Context, categoryId string) ([]dto.ProductDTO, error)
	GetProductById(ctx context.Context, id string) (dto.ProductDTO, error)
	CreateProduct(ctx context.Context, product dto.CreateProductDTO) (dto.ProductDTO, error)
	UpdateProduct(ctx context.Context, product dto.UpdateProductDTO) (dto.ProductDTO, error)
	DeleteProduct(ctx context.Context, id string) error
}
