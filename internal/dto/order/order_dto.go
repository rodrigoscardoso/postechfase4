package dto

import (
	entity "post-tech-challenge-10soat/internal/entities"
	"time"
)

type OrderStatus string

const (
	OrderStatusReceived  OrderStatus = "received"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusCompleted OrderStatus = "completed"
)

type OrderDTO struct {
	Id        string
	Number    int
	Status    string
	ClientId  string
	PaymentId string
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d OrderDTO) ToEntity() entity.Order {
	return entity.Order{
		Id:        d.Id,
		Number:    d.Number,
		Status:    entity.OrderStatus(d.Status),
		ClientId:  d.ClientId,
		PaymentId: d.PaymentId,
		Total:     d.Total,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
