package interfaces

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/category"
)

type CategoryRepository interface {
	GetCategoryById(ctx context.Context, categoryId string) (dto.CategoryDTO, error)
}
