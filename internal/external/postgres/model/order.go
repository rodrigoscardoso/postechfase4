package model

import (
	"database/sql"
	dto "post-tech-challenge-10soat/internal/dto/order"
	"time"
)

type OrderModel struct {
	Id        string         `db:"id"`
	Number    int            `db:"number"`
	Status    string         `db:"status"`
	ClientId  string         `db:"clientId"`
	PaymentId sql.NullString `db:"paymentId"`
	Total     float64        `db:"total"`
	CreatedAt time.Time      `db:"createdAt"`
	UpdatedAt time.Time      `db:"updatedAt"`
}

func (m OrderModel) ToDTO() dto.OrderDTO {
	return dto.OrderDTO{
		Id:        m.Id,
		Number:    m.Number,
		Status:    m.Status,
		ClientId:  m.ClientId,
		PaymentId: m.PaymentId.String,
		Total:     m.Total,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
