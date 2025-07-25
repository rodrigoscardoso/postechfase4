package client

import (
	"context"
	"fmt"
	dto "post-tech-challenge-10soat/internal/dto/client"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
)

type CreateClientUseCaseImpl struct {
	gateway interfaces.ClientGateway
}

func NewCreateClientUsecaseImpl(gateway interfaces.ClientGateway) CreateClientUseCase {
	return &CreateClientUseCaseImpl{
		gateway,
	}
}

func (s CreateClientUseCaseImpl) Execute(ctx context.Context, createClientDTO dto.CreateClientDTO) (entity.Client, error) {
	newClient := entity.Client{
		Cpf:   createClientDTO.Cpf,
		Name:  createClientDTO.Name,
		Email: createClientDTO.Email,
	}
	client, err := s.gateway.CreateClient(ctx, newClient)
	if err != nil {
		return entity.Client{}, fmt.Errorf("failed to create client - %s", err.Error())
	}
	return client, nil
}
