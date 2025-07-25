package mapper

import (
	entity "post-tech-challenge-10soat/internal/entities"
	"post-tech-challenge-10soat/internal/utils"

	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID   uuid.UUID `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Name string    `json:"name" example:"Lanche"`
}

func NewCategoryResponse(category entity.Category) CategoryResponse {
	return CategoryResponse{
		ID:   utils.StringToUuid(category.Id),
		Name: category.Name,
	}
}
