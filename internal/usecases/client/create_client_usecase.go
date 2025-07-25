package client

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/client"
	entity "post-tech-challenge-10soat/internal/entities"
)

type CreateClientUseCase interface {
	Execute(ctx context.Context, client dto.CreateClientDTO) (entity.Client, error)
}
