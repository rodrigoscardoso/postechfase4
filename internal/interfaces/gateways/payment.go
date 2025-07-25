package interfaces

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type PaymentGateway interface {
	CreatePayment(ctx context.Context, payment entity.Payment) (entity.Payment, error)
}
