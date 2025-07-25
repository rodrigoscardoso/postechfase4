package client

import (
	"context"
	"fmt"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
)

type GetClientByIdUseCaseImpl struct {
	gateway interfaces.ClientGateway
}

func NewGetClientByIdUseCaseImpl(gateway interfaces.ClientGateway) GetClientByIdUseCase {
	return &GetClientByIdUseCaseImpl{
		gateway,
	}
}

func (s GetClientByIdUseCaseImpl) Execute(ctx context.Context, id string) (entity.Client, error) {
	client, err := s.gateway.GetClientById(ctx, id)
	if err != nil {
		return entity.Client{}, fmt.Errorf("failed to get client by id - %s", err.Error())
	}
	return client, nil
}
