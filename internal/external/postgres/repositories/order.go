package repository

import (
	"context"
	"fmt"
	dto "post-tech-challenge-10soat/internal/dto/order"
	"post-tech-challenge-10soat/internal/external/postgres"
	"post-tech-challenge-10soat/internal/external/postgres/model"

	sq "github.com/Masterminds/squirrel"
)

type OrderRepositoryImpl struct {
	db *postgres.DB
}

func NewOrderRepositoryImpl(db *postgres.DB) OrderRepositoryImpl {
	return OrderRepositoryImpl{
		db,
	}
}

func (repository OrderRepositoryImpl) CreateOrder(ctx context.Context, order dto.CreateOrderDTO) (dto.OrderDTO, error) {
	var orderModel model.OrderModel
	query := repository.db.QueryBuilder.Insert("orders").
		Columns("status", "client_id", "total").
		Values(order.Status, order.ClientId, order.Total).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return dto.OrderDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&orderModel.Id,
		&orderModel.Number,
		&orderModel.Status,
		&orderModel.ClientId,
		&orderModel.PaymentId,
		&orderModel.Total,
		&orderModel.CreatedAt,
		&orderModel.UpdatedAt,
	)
	if err != nil {
		return dto.OrderDTO{}, err
	}
	return orderModel.ToDTO(), nil
}

func (repository OrderRepositoryImpl) DeleteOrder(ctx context.Context, id string) error {
	query := repository.db.QueryBuilder.Delete("orders").
		Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = repository.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (repository OrderRepositoryImpl) ListOrders(ctx context.Context, limit uint64) ([]dto.OrderDTO, error) {
	var orderModel model.OrderModel
	var orders []dto.OrderDTO
	query := repository.db.QueryBuilder.Select("*").
		From("orders").
		Where(sq.Or{
			sq.Eq{"status": "received"},
			sq.Eq{"status": "preparing"},
			sq.Eq{"status": "ready"},
		}).
		OrderBy("created_at ASC").
		Limit(limit)
	sql, args, err := query.ToSql()
	if err != nil {
		return []dto.OrderDTO{}, fmt.Errorf("failed to get orders - %s", err.Error())
	}
	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return []dto.OrderDTO{}, fmt.Errorf("failed to get orders - %s", err.Error())
	}
	for rows.Next() {
		err := rows.Scan(
			&orderModel.Id,
			&orderModel.Number,
			&orderModel.Status,
			&orderModel.ClientId,
			&orderModel.PaymentId,
			&orderModel.Total,
			&orderModel.CreatedAt,
			&orderModel.UpdatedAt,
		)
		if err == nil {
			order := orderModel.ToDTO()
			orders = append(orders, order)
		}
	}
	return orders, nil
}

func (repository OrderRepositoryImpl) GetOrderById(ctx context.Context, id string) (dto.OrderDTO, error) {
	var orderModel model.OrderModel
	query := repository.db.QueryBuilder.Select("*").
		From("orders").
		Where(sq.Eq{"id": id}).
		Limit(1)
	sql, args, err := query.ToSql()
	if err != nil {
		return dto.OrderDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&orderModel.Id,
		&orderModel.Number,
		&orderModel.Status,
		&orderModel.ClientId,
		&orderModel.PaymentId,
		&orderModel.Total,
		&orderModel.CreatedAt,
		&orderModel.UpdatedAt,
	)
	if err != nil {
		return dto.OrderDTO{}, err
	}
	return orderModel.ToDTO(), nil
}

func (repository OrderRepositoryImpl) UpdateOrderStatus(ctx context.Context, id string, status string) (dto.OrderDTO, error) {
	var orderModel model.OrderModel
	query := repository.db.QueryBuilder.Update("orders").
		Set("status", sq.Expr("COALESCE(?, status)", status)).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return dto.OrderDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&orderModel.Id,
		&orderModel.Number,
		&orderModel.Status,
		&orderModel.ClientId,
		&orderModel.PaymentId,
		&orderModel.Total,
		&orderModel.CreatedAt,
		&orderModel.UpdatedAt,
	)
	if err != nil {
		return dto.OrderDTO{}, err
	}
	return orderModel.ToDTO(), nil
}
