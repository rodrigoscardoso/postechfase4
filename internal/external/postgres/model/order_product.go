package model

import (
	dto "post-tech-challenge-10soat/internal/dto/order"
	"time"
)

type OrderProductModel struct {
	Id          string    `db:"id"`
	OrderId     string    `db:"orderId"`
	ProductId   string    `db:"productId"`
	Quantity    int       `db:"quantity"`
	SubTotal    float64   `db:"subTotal"`
	Observation string    `db:"observation"`
	CreatedAt   time.Time `db:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt"`
}

func (m OrderProductModel) ToDTO() dto.OrderProductDTO {
	return dto.OrderProductDTO{
		Id:          m.Id,
		OrderId:     m.OrderId,
		ProductId:   m.ProductId,
		Quantity:    m.Quantity,
		SubTotal:    m.SubTotal,
		Observation: m.Observation,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
