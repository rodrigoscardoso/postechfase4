package interfaces

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/order"
)

type OrderProductRepository interface {
	CreateOrderProduct(ctx context.Context, orderProduct dto.CreateOrderProductDTO) (dto.OrderProductDTO, error)
}
