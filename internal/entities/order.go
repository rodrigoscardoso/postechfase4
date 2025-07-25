package entity

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusPaymentPending OrderStatus = "payment_pending"
	OrderStatusReceived       OrderStatus = "received"
	OrderStatusPreparing      OrderStatus = "preparing"
	OrderStatusReady          OrderStatus = "ready"
	OrderStatusCompleted      OrderStatus = "completed"
)

type Order struct {
	Id        string
	Number    int
	Status    OrderStatus
	ClientId  string
	PaymentId string
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
