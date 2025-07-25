package payment

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/payment"
	entity "post-tech-challenge-10soat/internal/entities"
)

type PaymentCheckoutUseCase interface {
	Execute(ctx context.Context, createPayment dto.CreatePaymentDTO) (entity.Payment, error)
}
