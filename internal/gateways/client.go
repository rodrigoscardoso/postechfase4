package gateways

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/client"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/repositories"
)

type ClientGatewayImpl struct {
	repository interfaces.ClientRepository
}

func NewClientGatewayImpl(repository interfaces.ClientRepository) *ClientGatewayImpl {
	return &ClientGatewayImpl{
		repository,
	}
}

func (cg ClientGatewayImpl) CreateClient(ctx context.Context, client entity.Client) (entity.Client, error) {
	createClientDTO := dto.CreateClientDTO{
		Cpf:   client.Cpf,
		Name:  client.Name,
		Email: client.Email,
	}
	createdClient, err := cg.repository.CreateClient(ctx, createClientDTO)
	if err != nil {
		return entity.Client{}, err
	}
	return createdClient.ToEntity(), nil
}

func (cg ClientGatewayImpl) GetClientByCpf(ctx context.Context, cpf string) (entity.Client, error) {
	client, err := cg.repository.GetClientByCpf(ctx, cpf)
	if err != nil {
		return entity.Client{}, err
	}
	return client.ToEntity(), nil
}

func (cg ClientGatewayImpl) GetClientById(ctx context.Context, id string) (entity.Client, error) {
	client, err := cg.repository.GetClientById(ctx, id)
	if err != nil {
		return entity.Client{}, err
	}
	return client.ToEntity(), nil
}
