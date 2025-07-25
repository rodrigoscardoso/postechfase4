package gateways

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/order"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/repositories"
)

type OrderProductGatewayImpl struct {
	repository interfaces.OrderProductRepository
}

func NewOrderProductGatewayImpl(repository interfaces.OrderProductRepository) *OrderProductGatewayImpl {
	return &OrderProductGatewayImpl{
		repository,
	}
}

func (og OrderProductGatewayImpl) CreateOrderProduct(ctx context.Context, orderProduct entity.OrderProduct) (entity.OrderProduct, error) {
	orderProductDTO := dto.CreateOrderProductDTO{
		OrderId:     orderProduct.OrderId,
		ProductId:   orderProduct.ProductId,
		Quantity:    orderProduct.Quantity,
		SubTotal:    orderProduct.SubTotal,
		Observation: orderProduct.Observation,
	}
	createdOrderProduct, err := og.repository.CreateOrderProduct(ctx, orderProductDTO)
	if err != nil {
		return entity.OrderProduct{}, nil
	}
	return createdOrderProduct.ToEntity(), nil
}
