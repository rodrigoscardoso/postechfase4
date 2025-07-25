package interfaces

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/payment"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment dto.CreatePaymentDTO) (dto.PaymentDTO, error)
}
