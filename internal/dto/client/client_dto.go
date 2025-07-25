package dto

import (
	entity "post-tech-challenge-10soat/internal/entities"
	"time"
)

type ClientDTO struct {
	Id        string
	Cpf       string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d ClientDTO) ToEntity() entity.Client {
	return entity.Client{
		Id:        d.Id,
		Cpf:       d.Cpf,
		Name:      d.Name,
		Email:     d.Email,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func (d ClientDTO) FromEntity(client entity.Client) ClientDTO {
	return ClientDTO{
		Id:        client.Id,
		Cpf:       client.Cpf,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}
}
