package model

import (
	"database/sql"
	dto "post-tech-challenge-10soat/internal/dto/client"
	"time"
)

type ClientModel struct {
	Id        string         `db:"id"`
	Cpf       sql.NullString `db:"cpf"`
	Name      string         `db:"name"`
	Email     string         `db:"email"`
	CreatedAt time.Time      `db:"createdAt"`
	UpdatedAt time.Time      `db:"updatedAt"`
}

func (m ClientModel) ToDTO() dto.ClientDTO {
	return dto.ClientDTO{
		Id:        m.Id,
		Cpf:       m.Cpf.String,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
