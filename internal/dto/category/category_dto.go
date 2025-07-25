package dto

import (
	entity "post-tech-challenge-10soat/internal/entities"
	"time"
)

type CategoryDTO struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d CategoryDTO) ToEntity() entity.Category {
	return entity.Category{
		Id:        d.Id,
		Name:      d.Name,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
