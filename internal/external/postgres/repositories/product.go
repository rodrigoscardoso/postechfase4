package repository

import (
	"context"
	"fmt"
	dto "post-tech-challenge-10soat/internal/dto/product"
	"post-tech-challenge-10soat/internal/external/postgres"
	"post-tech-challenge-10soat/internal/external/postgres/model"

	"post-tech-challenge-10soat/internal/utils"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type ProductRepositoryImpl struct {
	db *postgres.DB
}

func NewProductRepositoryImpl(db *postgres.DB) ProductRepositoryImpl {
	return ProductRepositoryImpl{
		db,
	}
}

func (repository ProductRepositoryImpl) ListProducts(ctx context.Context, categoryId string) ([]dto.ProductDTO, error) {
	var productModel model.ProductModel
	var products []dto.ProductDTO
	query := repository.db.QueryBuilder.Select("*").
		From("products").
		OrderBy("created_at")

	if categoryId != "" {
		err := uuid.Validate(categoryId)
		if err != nil {
			return []dto.ProductDTO{}, fmt.Errorf("invalid category")
		}
		query = query.Where(sq.Eq{"category_id": categoryId})
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return []dto.ProductDTO{}, err
	}
	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return []dto.ProductDTO{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&productModel.Id,
			&productModel.Name,
			&productModel.Description,
			&productModel.Image,
			&productModel.Value,
			&productModel.CategoryId,
			&productModel.CreatedAt,
			&productModel.UpdatedAt,
		)
		if err != nil {
			return []dto.ProductDTO{}, err
		}
		product := productModel.ToDTO()
		products = append(products, product)
	}
	return products, nil
}

func (repository ProductRepositoryImpl) GetProductById(ctx context.Context, id string) (dto.ProductDTO, error) {
	var productModel model.ProductModel
	query := repository.db.QueryBuilder.Select("*").
		From("products").
		Where(sq.Eq{"id": id}).
		Limit(1)
	sql, args, err := query.ToSql()
	if err != nil {
		return dto.ProductDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&productModel.Id,
		&productModel.Name,
		&productModel.Description,
		&productModel.Image,
		&productModel.Value,
		&productModel.CategoryId,
		&productModel.CreatedAt,
		&productModel.UpdatedAt,
	)
	if err != nil {
		return dto.ProductDTO{}, err
	}
	return productModel.ToDTO(), nil
}

func (repository ProductRepositoryImpl) CreateProduct(ctx context.Context, product dto.CreateProductDTO) (dto.ProductDTO, error) {
	var productModel model.ProductModel
	query := repository.db.QueryBuilder.Insert("products").
		Columns("name", "description", "image", "value", "category_id").
		Values(product.Name, product.Description, product.Image, product.Value, product.CategoryId).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return dto.ProductDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&productModel.Id,
		&productModel.Name,
		&productModel.Description,
		&productModel.Image,
		&productModel.Value,
		&productModel.CategoryId,
		&productModel.CreatedAt,
		&productModel.UpdatedAt,
	)
	if err != nil {
		return dto.ProductDTO{}, err
	}
	return productModel.ToDTO(), nil
}

func (repository ProductRepositoryImpl) UpdateProduct(ctx context.Context, product dto.UpdateProductDTO) (dto.ProductDTO, error) {
	var productModel model.ProductModel
	name := utils.NullString(product.Name)
	description := utils.NullString(product.Description)
	image := utils.NullString(product.Image)
	query := repository.db.QueryBuilder.Update("products").
		Set("name", sq.Expr("COALESCE(?, name)", name)).
		Set("description", sq.Expr("COALESCE(?, description)", description)).
		Set("image", sq.Expr("COALESCE(?, image)", image)).
		Set("value", sq.Expr("COALESCE(?, value)", product.Value)).
		Set("category_id", sq.Expr("COALESCE(?, category_id)", product.CategoryId)).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": product.Id}).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return dto.ProductDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&productModel.Id,
		&productModel.Name,
		&productModel.Description,
		&productModel.Image,
		&productModel.Value,
		&productModel.CategoryId,
		&productModel.CreatedAt,
		&productModel.UpdatedAt,
	)
	if err != nil {
		return dto.ProductDTO{}, err
	}
	return productModel.ToDTO(), nil
}

func (repository ProductRepositoryImpl) DeleteProduct(ctx context.Context, id string) error {
	query := repository.db.QueryBuilder.Delete("products").
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
