package controllers

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/client"
	entity "post-tech-challenge-10soat/internal/entities"
	"post-tech-challenge-10soat/internal/usecases/client"
)

// ClientController defines the interface for client controller
type ClientController interface {
	CreateClient(ctx context.Context, createClient dto.CreateClientDTO) (entity.Client, error)
	GetClientByCpf(ctx context.Context, cpf string) (entity.Client, error)
	GetClientById(ctx context.Context, id string) (entity.Client, error)
}

type clientController struct {
	getClientByCpf client.GetClientByCpfUseCase
	getClientById  client.GetClientByIdUseCase
	createClient   client.CreateClientUseCase
}

func NewClientController(
	getClientByCpf client.GetClientByCpfUseCase,
	getClientById client.GetClientByIdUseCase,
	createClient client.CreateClientUseCase,
) ClientController {
	return &clientController{
		getClientByCpf: getClientByCpf,
		getClientById:  getClientById,
		createClient:   createClient,
	}
}

func (c *clientController) GetClientByCpf(ctx context.Context, cpf string) (entity.Client, error) {
	client, err := c.getClientByCpf.Execute(ctx, cpf)
	if err != nil {
		return entity.Client{}, err
	}
	return client, nil
}

func (c *clientController) GetClientById(ctx context.Context, id string) (entity.Client, error) {
	client, err := c.getClientById.Execute(ctx, id)
	if err != nil {
		return entity.Client{}, err
	}
	return client, nil
}

func (c *clientController) CreateClient(ctx context.Context, createClient dto.CreateClientDTO) (entity.Client, error) {
	client, err := c.createClient.Execute(ctx, createClient)
	if err != nil {
		return entity.Client{}, err
	}
	return client, nil
}
