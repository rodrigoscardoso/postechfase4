package interfaces

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type ClientGateway interface {
	CreateClient(ctx context.Context, client entity.Client) (entity.Client, error)
	GetClientByCpf(ctx context.Context, cpf string) (entity.Client, error)
	GetClientById(ctx context.Context, id string) (entity.Client, error)
}
