package repository

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/payment"
	"post-tech-challenge-10soat/internal/external/postgres"
	"post-tech-challenge-10soat/internal/external/postgres/model"
)

type PaymentRepositoryImpl struct {
	db *postgres.DB
}

func NewPaymentRepositoryImpl(db *postgres.DB) PaymentRepositoryImpl {
	return PaymentRepositoryImpl{
		db,
	}
}

func (repository PaymentRepositoryImpl) CreatePayment(ctx context.Context, payment dto.CreatePaymentDTO) (dto.PaymentDTO, error) {
	var paymentModel model.PaymentModel
	query := repository.db.QueryBuilder.Insert("payments").
		Columns("type", "provider").
		Values(payment.Type, payment.Provider).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return dto.PaymentDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&paymentModel.Id,
		&paymentModel.Type,
		&paymentModel.Provider,
		&paymentModel.CreatedAt,
		&paymentModel.UpdatedAt,
	)
	if err != nil {
		return dto.PaymentDTO{}, err
	}
	return paymentModel.ToDTO(), nil
}
