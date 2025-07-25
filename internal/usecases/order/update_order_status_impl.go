package order

import (
	"context"
	"errors"
	"fmt"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
	"post-tech-challenge-10soat/internal/utils"
)

type UpdateOrderStatusUseCaseImpl struct {
	orderGateway interfaces.OrderGateway
}

func NewUpdateOrderStatusUseCaseImpl(orderGateway interfaces.OrderGateway) UpdateOrderStatusUseCase {
	return &UpdateOrderStatusUseCaseImpl{
		orderGateway,
	}
}

func (u UpdateOrderStatusUseCaseImpl) Execute(ctx context.Context, id string, status string) (entity.Order, error) {
	order, err := u.orderGateway.GetOrderById(ctx, id)
	if err != nil {
		return entity.Order{}, err
	}
	validTransitions := map[string][]string{
		"received":  {"preparing"},
		"preparing": {"ready"},
		"ready":     {"completed"},
		"completed": {},
	}
	allowed, exists := validTransitions[string(order.Status)]
	if !exists {
		return entity.Order{}, errors.New("invalid status")
	}
	if !utils.Contains(allowed, status) {
		return entity.Order{}, fmt.Errorf("cannot update order status for '%s' to '%s'", order.Status, status)
	}
	updatedOrder, err := u.orderGateway.UpdateOrderStatus(ctx, id, status)
	if err != nil {
		return entity.Order{}, err
	}
	return updatedOrder, nil
}
