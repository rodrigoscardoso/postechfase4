package interfaces

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/client"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, client dto.CreateClientDTO) (dto.ClientDTO, error)
	GetClientByCpf(ctx context.Context, cpf string) (dto.ClientDTO, error)
	GetClientById(ctx context.Context, id string) (dto.ClientDTO, error)
}
