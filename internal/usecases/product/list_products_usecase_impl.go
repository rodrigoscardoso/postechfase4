package product

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
)

type ListProductsUseCaseImpl struct {
	productGateway  interfaces.ProductGateway
	categoryGateway interfaces.CategoryGateway
}

func NewListProductsUsecaseImpl(productGateway interfaces.ProductGateway, categoryGateway interfaces.CategoryGateway) ListProductsUseCase {
	return &ListProductsUseCaseImpl{
		productGateway,
		categoryGateway,
	}
}

func (s ListProductsUseCaseImpl) Execute(ctx context.Context, categoryId string) ([]entity.Product, error) {
	var products []entity.Product
	products, err := s.productGateway.ListProducts(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	for i, product := range products {
		category, err := s.categoryGateway.GetCategoryById(ctx, product.CategoryId)
		if err != nil {
			if err == entity.ErrDataNotFound {
				return nil, err
			}
			return nil, err
		}

		products[i].Category = category
	}
	return products, nil
}
