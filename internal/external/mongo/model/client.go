package model

import (
	dto "post-tech-challenge-10soat/internal/dto/client"
	"time"
)

type ClientModel struct {
	Id        string    `bson:"_id,omitempty"`
	Cpf       string    `bson:"cpf"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (m ClientModel) ToDTO() dto.ClientDTO {
	return dto.ClientDTO{
		Id:        m.Id,
		Cpf:       m.Cpf,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
