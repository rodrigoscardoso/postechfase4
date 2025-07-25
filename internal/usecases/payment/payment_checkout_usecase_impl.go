package payment

import (
	"context"
	"fmt"
	dto "post-tech-challenge-10soat/internal/dto/payment"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
)

type PaymentCheckoutUseCaseImpl struct {
	gateway interfaces.PaymentGateway
}

func NewPaymentCheckoutUsecaseImpl(gateway interfaces.PaymentGateway) PaymentCheckoutUseCase {
	return &PaymentCheckoutUseCaseImpl{
		gateway,
	}
}

func (s PaymentCheckoutUseCaseImpl) Execute(ctx context.Context, createPayment dto.CreatePaymentDTO) (entity.Payment, error) {
	//TODO: Checkout no fornecedor MP
	paymentInfo := entity.Payment{
		Type:     createPayment.Type,
		Provider: createPayment.Provider,
	}
	payment, err := s.gateway.CreatePayment(ctx, paymentInfo)
	if err != nil {
		if err == entity.ErrConflictingData {
			return entity.Payment{}, err
		}
		return entity.Payment{}, fmt.Errorf("failed to make payment - %s", err.Error())
	}
	return payment, nil
}
