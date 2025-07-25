package model

import (
	dto "post-tech-challenge-10soat/internal/dto/payment"
	"time"
)

type PaymentModel struct {
	Id        string    `db:"id"`
	Provider  string    `db:"provider"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

func (m PaymentModel) ToDTO() dto.PaymentDTO {
	return dto.PaymentDTO{
		Id:        m.Id,
		Provider:  m.Provider,
		Type:      m.Type,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
