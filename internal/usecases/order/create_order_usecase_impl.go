package order

import (
	"context"
	"fmt"
	dto "post-tech-challenge-10soat/internal/dto/order"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"

	"github.com/google/uuid"
)

type CreateOrderUsecaseImpl struct {
	productGateway      interfaces.ProductGateway
	clientGateway       interfaces.ClientGateway
	orderGateway        interfaces.OrderGateway
	orderProductGateway interfaces.OrderProductGateway
}

func NewCreateOrderUsecaseImpl(
	productGateway interfaces.ProductGateway,
	clientGateway interfaces.ClientGateway,
	orderGateway interfaces.OrderGateway,
	orderProductGateway interfaces.OrderProductGateway,
) CreateOrderUseCase {
	return &CreateOrderUsecaseImpl{
		productGateway,
		clientGateway,
		orderGateway,
		orderProductGateway,
	}
}

func (s CreateOrderUsecaseImpl) Execute(ctx context.Context, createOrder dto.CreateOrderDTO) (entity.Order, error) {
	var totalValue float64
	for _, orderProduct := range createOrder.Products {
		product, err := s.productGateway.GetProductById(ctx, orderProduct.ProductId)
		if err != nil {
			if err == entity.ErrDataNotFound {
				return entity.Order{}, err
			}
			return entity.Order{}, fmt.Errorf("cannot create order because has invalid product - %s", err.Error())
		}
		subTotal := product.Value * float64(orderProduct.Quantity)
		totalValue += subTotal
	}

	orderInfo := entity.Order{
		Status: entity.OrderStatusPaymentPending,
		Total:  totalValue,
	}
	if createOrder.ClientId != "" && uuid.Validate(createOrder.ClientId) == nil {
		client, err := s.clientGateway.GetClientById(ctx, createOrder.ClientId)
		if err != nil {
			if err == entity.ErrDataNotFound {
				return entity.Order{}, err
			}
			return entity.Order{}, fmt.Errorf("cannot create order because has invalid client - %s", err.Error())
		}
		orderInfo.ClientId = client.Id
	} else {
		orderInfo.ClientId = ""
	}
	order, err := s.orderGateway.CreateOrder(ctx, orderInfo)
	if err != nil {
		if err == entity.ErrDataNotFound {
			return entity.Order{}, err
		}
		return entity.Order{}, fmt.Errorf("cannot create order - %s", err.Error())
	}
	for _, orderProduct := range createOrder.Products {
		product, err := s.productGateway.GetProductById(ctx, orderProduct.ProductId)
		if err != nil {
			if err == entity.ErrDataNotFound {
				return entity.Order{}, err
			}
			return entity.Order{}, fmt.Errorf("cannot create order because has invalid product - %s", err.Error())
		}
		subTotal := product.Value * float64(orderProduct.Quantity)
		orderProductInfo := entity.OrderProduct{
			OrderId:     order.Id,
			ProductId:   product.Id,
			Quantity:    orderProduct.Quantity,
			SubTotal:    subTotal,
			Observation: orderProduct.Observation,
		}
		_, err = s.orderProductGateway.CreateOrderProduct(ctx, orderProductInfo)
		if err != nil {
			if err == entity.ErrDataNotFound {
				return entity.Order{}, err
			}
			err := s.orderGateway.DeleteOrder(ctx, order.Id)
			if err != nil {
				return entity.Order{}, err
			}
			return entity.Order{}, fmt.Errorf("cannot complete order - %s", err.Error())
		}
	}
	return order, nil
}
