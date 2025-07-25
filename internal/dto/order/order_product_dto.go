package dto

import (
	entity "post-tech-challenge-10soat/internal/entities"
	"time"
)

type OrderProductDTO struct {
	Id          string
	OrderId     string
	ProductId   string
	Quantity    int
	SubTotal    float64
	Observation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d OrderProductDTO) ToEntity() entity.OrderProduct {
	return entity.OrderProduct{
		Id:          d.Id,
		OrderId:     d.OrderId,
		ProductId:   d.ProductId,
		Quantity:    d.Quantity,
		SubTotal:    d.SubTotal,
		Observation: d.Observation,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
