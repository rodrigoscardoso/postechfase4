package client

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
)

type GetClientByIdUseCase interface {
	Execute(ctx context.Context, id string) (entity.Client, error)
}
