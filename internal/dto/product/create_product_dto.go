package dto

import dto "post-tech-challenge-10soat/internal/dto/category"

type CreateProductDTO struct {
	Name        string
	Description string
	Image       string
	Value       float64
	CategoryId  string
	CategoryDTO dto.CategoryDTO
}
