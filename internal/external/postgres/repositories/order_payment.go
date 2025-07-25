package repository

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/order"
	"post-tech-challenge-10soat/internal/external/postgres"
	"post-tech-challenge-10soat/internal/external/postgres/model"
)

type OrderProductRepositoryImpl struct {
	db *postgres.DB
}

func NewOrderProductRepositoryImpl(db *postgres.DB) OrderProductRepositoryImpl {
	return OrderProductRepositoryImpl{
		db,
	}
}

func (repository OrderProductRepositoryImpl) CreateOrderProduct(ctx context.Context, orderProduct dto.CreateOrderProductDTO) (dto.OrderProductDTO, error) {
	var orderProductModel model.OrderProductModel
	query := repository.db.QueryBuilder.Insert("order_products").
		Columns("order_id", "product_id", "quantity", "sub_total", "observation").
		Values(orderProduct.OrderId, orderProduct.ProductId, orderProduct.Quantity, orderProduct.SubTotal, orderProduct.Observation).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return dto.OrderProductDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&orderProductModel.Id,
		&orderProductModel.OrderId,
		&orderProductModel.ProductId,
		&orderProductModel.Quantity,
		&orderProductModel.SubTotal,
		&orderProductModel.Observation,
		&orderProductModel.CreatedAt,
		&orderProductModel.UpdatedAt,
	)
	if err != nil {
		return dto.OrderProductDTO{}, err
	}
	return orderProductModel.ToDTO(), nil
}
