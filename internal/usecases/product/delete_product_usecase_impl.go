package product

import (
	"context"
	"fmt"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"

	"github.com/google/uuid"
)

type DeleteProductUsecaseImpl struct {
	gateway interfaces.ProductGateway
}

func NewDeleteProductUsecaseImpl(gateway interfaces.ProductGateway) DeleteProductUseCase {
	return &DeleteProductUsecaseImpl{
		gateway,
	}
}

func (s DeleteProductUsecaseImpl) Execute(ctx context.Context, id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid product id")
	}
	_, err = s.gateway.GetProductById(ctx, id)
	if err != nil {
		if err == entity.ErrDataNotFound {
			return err
		}
		return fmt.Errorf("cannot delete product for this identifier - %s", err.Error())
	}
	return s.gateway.DeleteProduct(ctx, id)
}
