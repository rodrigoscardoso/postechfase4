package order

import (
	"context"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
	"sort"
)

type ListOrdersUseCaseImpl struct {
	orderGateway interfaces.OrderGateway
}

func NewListOrdersUseCaseImpl(orderGateway interfaces.OrderGateway) ListOrdersUseCase {
	return &ListOrdersUseCaseImpl{
		orderGateway,
	}
}

func (l ListOrdersUseCaseImpl) Execute(ctx context.Context, limit uint64) ([]entity.Order, error) {
	orders, err := l.orderGateway.ListOrders(ctx, limit)
	if err != nil {
		return []entity.Order{}, err
	}
	sortOrdersbyStatus(orders)
	return orders, nil
}

func sortOrdersbyStatus(orders []entity.Order) {
	statusPriority := map[string]int{
		"ready":     1,
		"preparing": 2,
		"received":  3,
	}
	sort.SliceStable(orders, func(i, j int) bool {
		if statusPriority[string(orders[i].Status)] != statusPriority[string(orders[j].Status)] {
			return statusPriority[string(orders[i].Status)] < statusPriority[string(orders[j].Status)]
		}
		return orders[i].CreatedAt.Before(orders[j].CreatedAt)
	})
}
