package mapper

import (
	entity "post-tech-challenge-10soat/internal/entities"
	"post-tech-challenge-10soat/internal/utils"
	"time"

	"github.com/google/uuid"
)

type ProductResponse struct {
	ID          uuid.UUID        `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Name        string           `json:"name" example:"Lanche 1"`
	Description string           `json:"description" example:"Lanche com bacon"`
	Image       string           `json:"image" example:"https://"`
	Value       float64          `json:"value" example:"10.90"`
	Category    CategoryResponse `json:"category"`
	CreatedAt   time.Time        `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt   time.Time        `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

func NewProductResponse(product entity.Product) ProductResponse {
	return ProductResponse{
		ID:          utils.StringToUuid(product.Id),
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		Value:       product.Value,
		Category:    NewCategoryResponse(product.Category),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
