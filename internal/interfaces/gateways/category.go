package interfaces

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type CategoryGateway interface {
	GetCategoryById(ctx context.Context, categoryId string) (entity.Category, error)
}
