package model

import (
	dto "post-tech-challenge-10soat/internal/dto/product"
	"time"
)

type ProductModel struct {
	Id            string        `db:"id"`
	Name          string        `db:"name"`
	Description   string        `db:"description"`
	Image         string        `db:"image"`
	Value         float64       `db:"value"`
	CategoryId    string        `db:"categoryId"`
	CategoryModel CategoryModel `db:"categoryModel"`
	CreatedAt     time.Time     `db:"createdAt"`
	UpdatedAt     time.Time     `db:"updatedAt"`
}

func (m ProductModel) ToDTO() dto.ProductDTO {
	return dto.ProductDTO{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		Image:       m.Image,
		Value:       m.Value,
		CategoryId:  m.CategoryId,
		CategoryDTO: m.CategoryModel.ToDTO(),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
