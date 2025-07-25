package product

import (
	"context"
	"fmt"
	dto "post-tech-challenge-10soat/internal/dto/product"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
)

type CreateProductUseCaseImpl struct {
	productGateway  interfaces.ProductGateway
	categoryGateway interfaces.CategoryGateway
}

func NewCreateProductUsecaseImpl(productGateway interfaces.ProductGateway, categoryGateway interfaces.CategoryGateway) CreateProductUseCase {
	return &CreateProductUseCaseImpl{
		productGateway,
		categoryGateway,
	}
}

func (s CreateProductUseCaseImpl) Execute(ctx context.Context, createProductDTO dto.CreateProductDTO) (entity.Product, error) {
	category, err := s.categoryGateway.GetCategoryById(ctx, createProductDTO.CategoryId)
	if err != nil {
		if err == entity.ErrDataNotFound {
			return entity.Product{}, err
		}
		return entity.Product{}, fmt.Errorf("cannot create product for this category - %s", err.Error())
	}
	newProduct := entity.Product{
		Name:        createProductDTO.Name,
		Description: createProductDTO.Description,
		Image:       createProductDTO.Image,
		Value:       createProductDTO.Value,
		CategoryId:  createProductDTO.CategoryId,
		Category:    category,
	}
	product, err := s.productGateway.CreateProduct(ctx, newProduct)
	if err != nil {
		if err == entity.ErrConflictingData {
			return entity.Product{}, err
		}
		return entity.Product{}, entity.ErrInternal
	}
	return product, nil
}
