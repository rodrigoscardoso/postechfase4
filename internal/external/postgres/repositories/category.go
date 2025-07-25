package repository

import (
	"context"
	"errors"
	dto "post-tech-challenge-10soat/internal/dto/category"
	"post-tech-challenge-10soat/internal/external/postgres"
	"post-tech-challenge-10soat/internal/external/postgres/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

type CategoryRepositoryImpl struct {
	db *postgres.DB
}

func NewCategoryRepositoryImpl(db *postgres.DB) CategoryRepositoryImpl {
	return CategoryRepositoryImpl{
		db,
	}
}

func (cr CategoryRepositoryImpl) GetCategoryById(ctx context.Context, id string) (dto.CategoryDTO, error) {
	var categoryModel model.CategoryModel
	query := cr.db.QueryBuilder.Select("*").
		From("categories").
		Where(sq.Eq{"id": id}).
		Limit(1)
	sql, args, err := query.ToSql()

	if err != nil {
		return dto.CategoryDTO{}, err
	}
	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&categoryModel.Id,
		&categoryModel.Name,
		&categoryModel.CreatedAt,
		&categoryModel.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return dto.CategoryDTO{}, errors.New("No data found")
		}
		return dto.CategoryDTO{}, err
	}
	return categoryModel.ToDTO(), nil
}
