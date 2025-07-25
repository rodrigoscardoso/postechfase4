package gateways

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/product"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/repositories"
)

type ProductGatewayImpl struct {
	repository interfaces.ProductRepository
}

func NewProductGatewayImpl(repository interfaces.ProductRepository) *ProductGatewayImpl {
	return &ProductGatewayImpl{
		repository,
	}
}

func (pg ProductGatewayImpl) ListProducts(ctx context.Context, categoryId string) ([]entity.Product, error) {
	var productsRes []entity.Product
	products, err := pg.repository.ListProducts(ctx, categoryId)
	if err != nil {
		return []entity.Product{}, err
	}
	for _, product := range products {
		productsRes = append(productsRes, product.ToEntity())
	}
	return productsRes, nil
}

func (pg ProductGatewayImpl) GetProductById(ctx context.Context, id string) (entity.Product, error) {
	product, err := pg.repository.GetProductById(ctx, id)
	if err != nil {
		return entity.Product{}, err
	}
	return product.ToEntity(), nil
}

func (pg ProductGatewayImpl) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	createProductDTO := dto.CreateProductDTO{
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		Value:       product.Value,
		CategoryId:  product.CategoryId,
	}
	createdProduct, err := pg.repository.CreateProduct(ctx, createProductDTO)
	if err != nil {
		return entity.Product{}, err
	}
	return createdProduct.ToEntity(), nil
}

func (pg ProductGatewayImpl) UpdateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	updateProductDTO := dto.UpdateProductDTO{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		Value:       product.Value,
		CategoryId:  product.CategoryId,
	}
	updatedProduct, err := pg.repository.UpdateProduct(ctx, updateProductDTO)
	if err != nil {
		return entity.Product{}, err
	}
	return updatedProduct.ToEntity(), nil
}

func (pg ProductGatewayImpl) DeleteProduct(ctx context.Context, id string) error {
	err := pg.repository.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
