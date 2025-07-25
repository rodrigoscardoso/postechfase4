package client

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type GetClientByCpfUseCase interface {
	Execute(ctx context.Context, cpf string) (entity.Client, error)
}
