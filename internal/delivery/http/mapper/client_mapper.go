package mapper

import (
	entity "post-tech-challenge-10soat/internal/entities"
	"post-tech-challenge-10soat/internal/utils"

	"github.com/google/uuid"
)

type ClientResponse struct {
	ID    uuid.UUID `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Name  string    `json:"name" example:"John Doe"`
	Email string    `json:"email" example:"john-doe@email.com"`
}

func NewClientResponse(client entity.Client) ClientResponse {
	return ClientResponse{
		ID:    utils.StringToUuid(client.Id),
		Name:  client.Name,
		Email: client.Email,
	}
}
