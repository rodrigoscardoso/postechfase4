package controllers

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/product"
	entity "post-tech-challenge-10soat/internal/entities"
	"post-tech-challenge-10soat/internal/usecases/product"
)

type ProductController struct {
	createProduct product.CreateProductUseCase
	deleteProduct product.DeleteProductUseCase
	updateProduct product.UpdateProductUseCase
	listProducts  product.ListProductsUseCase
}

func NewProductController(
	createProduct product.CreateProductUseCase,
	deleteProduct product.DeleteProductUseCase,
	updateProduct product.UpdateProductUseCase,
	listProducts product.ListProductsUseCase,
) *ProductController {
	return &ProductController{
		createProduct,
		deleteProduct,
		updateProduct,
		listProducts,
	}
}

func (c *ProductController) CreateProduct(ctx context.Context, createProductDTO dto.CreateProductDTO) (entity.Product, error) {
	product, err := c.createProduct.Execute(ctx, createProductDTO)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, id string) error {
	err := c.deleteProduct.Execute(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, updateProductDTO dto.UpdateProductDTO) (entity.Product, error) {
	product, err := c.updateProduct.Execute(ctx, updateProductDTO)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (c *ProductController) ListProducts(ctx context.Context, categoryId string) ([]entity.Product, error) {
	products, err := c.listProducts.Execute(ctx, categoryId)
	if err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}
