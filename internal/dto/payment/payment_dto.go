package dto

import (
	entity "post-tech-challenge-10soat/internal/entities"
	"time"
)

type PaymentDTO struct {
	Id        string
	Provider  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d PaymentDTO) ToEntity() entity.Payment {
	return entity.Payment{
		Id:        d.Id,
		Provider:  d.Provider,
		Type:      d.Type,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
