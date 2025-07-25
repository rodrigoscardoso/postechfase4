package category

import (
	"context"
	"fmt"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
)

type GetCategoryUsecaseImpl struct {
	gateway interfaces.CategoryGateway
}

func NewGetCategoryUsecase(gateway interfaces.CategoryGateway) GetCategoryUseCase {
	return &GetCategoryUsecaseImpl{
		gateway,
	}
}

func (s *GetCategoryUsecaseImpl) Execute(ctx context.Context, id string) (entity.Category, error) {
	category, err := s.gateway.GetCategoryById(ctx, id)
	if err != nil {
		return entity.Category{}, fmt.Errorf("failed to get category by id - %s", err.Error())
	}
	return category, nil
}
