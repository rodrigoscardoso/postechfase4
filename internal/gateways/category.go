package gateways

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/repositories"
)

type CategoryGatewayImpl struct {
	repository interfaces.CategoryRepository
}

func NewCategoryGatewayImpl(repository interfaces.CategoryRepository) *CategoryGatewayImpl {
	return &CategoryGatewayImpl{
		repository,
	}
}

func (cg CategoryGatewayImpl) GetCategoryById(ctx context.Context, categoryId string) (entity.Category, error) {
	category, err := cg.repository.GetCategoryById(ctx, categoryId)
	if err != nil {
		return entity.Category{}, err
	}
	return category.ToEntity(), nil
}
