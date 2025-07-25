package product

import "context"

type DeleteProductUseCase interface {
	Execute(ctx context.Context, id string) error
}
