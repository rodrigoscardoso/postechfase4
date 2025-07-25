package gateways

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/payment"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/repositories"
)

type PaymentGatewayImpl struct {
	repository interfaces.PaymentRepository
}

func NewPaymentGatewayImpl(repository interfaces.PaymentRepository) *PaymentGatewayImpl {
	return &PaymentGatewayImpl{
		repository,
	}
}

func (pg PaymentGatewayImpl) CreatePayment(ctx context.Context, payment entity.Payment) (entity.Payment, error) {
	createPaymentDTO := dto.CreatePaymentDTO{
		Provider: payment.Provider,
		Type:     payment.Type,
	}
	createdPayment, err := pg.repository.CreatePayment(ctx, createPaymentDTO)
	if err != nil {
		return entity.Payment{}, nil
	}
	return createdPayment.ToEntity(), nil
}
