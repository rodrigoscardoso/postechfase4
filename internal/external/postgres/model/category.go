package model

import (
	dto "post-tech-challenge-10soat/internal/dto/category"
	"time"
)

type CategoryModel struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

func (m CategoryModel) ToDTO() dto.CategoryDTO {
	return dto.CategoryDTO{
		Id:        m.Id,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
