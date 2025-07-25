package client

import (
	"context"
	"fmt"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
)

type GetClientByCpfUseCaseImpl struct {
	gateway interfaces.ClientGateway
}

func NewGetClientByCpfUseCaseImpl(gateway interfaces.ClientGateway) GetClientByCpfUseCase {
	return &GetClientByCpfUseCaseImpl{
		gateway,
	}
}

func (s GetClientByCpfUseCaseImpl) Execute(ctx context.Context, cpf string) (entity.Client, error) {
	client, err := s.gateway.GetClientByCpf(ctx, cpf)
	if err != nil {
		return entity.Client{}, fmt.Errorf("failed to get client by cpf - %s", err.Error())
	}
	return client, nil
}
