package controllers

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/order"
	entity "post-tech-challenge-10soat/internal/entities"
	"post-tech-challenge-10soat/internal/usecases/order"
)

type OrderController struct {
	createOrder           order.CreateOrderUseCase
	listOrders            order.ListOrdersUseCase
	getOrderPaymentStatus order.GetOrderPaymentStatusUseCase
	updateOrderStatus     order.UpdateOrderStatusUseCase
}

func NewOrderController(
	createOrder order.CreateOrderUseCase,
	listOrders order.ListOrdersUseCase,
	getOrderPaymentStatus order.GetOrderPaymentStatusUseCase,
	updateOrderStatus order.UpdateOrderStatusUseCase,
) *OrderController {
	return &OrderController{
		createOrder,
		listOrders,
		getOrderPaymentStatus,
		updateOrderStatus,
	}
}

func (c *OrderController) CreateOrder(ctx context.Context, createOrderDTO dto.CreateOrderDTO) (entity.Order, error) {
	order, err := c.createOrder.Execute(ctx, createOrderDTO)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (c *OrderController) ListOrders(ctx context.Context, limit uint64) ([]entity.Order, error) {
	orders, err := c.listOrders.Execute(ctx, limit)
	if err != nil {
		return []entity.Order{}, err
	}
	return orders, nil
}

func (c *OrderController) GetOrderPaymentStatus(ctx context.Context, id string) (order.OrderPaymentStatus, error) {
	orderPaymentStatus, err := c.getOrderPaymentStatus.Execute(ctx, id)
	if err != nil {
		return order.OrderPaymentStatus{}, err
	}
	return orderPaymentStatus, nil
}

func (c *OrderController) UpdateOrderStatus(ctx context.Context, id string, status string) (entity.Order, error) {
	order, err := c.updateOrderStatus.Execute(ctx, id, status)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
