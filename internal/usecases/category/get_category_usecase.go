package category

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type GetCategoryUseCase interface {
	Execute(ctx context.Context, id string) (entity.Category, error)
}
