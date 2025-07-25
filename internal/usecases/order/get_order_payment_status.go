package order

import (
	"context"
)

type PaymentStatus string

const (
	PaymentPending  PaymentStatus = "payment_pending"
	PaymentApproved PaymentStatus = "payment_approved"
)

type OrderPaymentStatus struct {
	PaymentStatus PaymentStatus
}

type GetOrderPaymentStatusUseCase interface {
	Execute(ctx context.Context, id string) (OrderPaymentStatus, error)
}
