package product

import (
	"context"
	"fmt"
	dto "post-tech-challenge-10soat/internal/dto/product"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"

	"github.com/google/uuid"
)

type UpdateProductUsecaseImpl struct {
	productGateway  interfaces.ProductGateway
	categoryGateway interfaces.CategoryGateway
}

func NewUpdateProductUsecaseImpl(productGateway interfaces.ProductGateway, categoryGateway interfaces.CategoryGateway) UpdateProductUseCase {
	return &UpdateProductUsecaseImpl{
		productGateway,
		categoryGateway,
	}
}

func (s UpdateProductUsecaseImpl) Execute(ctx context.Context, updateProductDTO dto.UpdateProductDTO) (entity.Product, error) {
	existingProduct, err := s.productGateway.GetProductById(ctx, updateProductDTO.Id)
	if err != nil {
		if err == entity.ErrDataNotFound {
			return entity.Product{}, err
		}
		return entity.Product{}, fmt.Errorf("cannot find product to update - %s", err.Error())
	}
	emptyData := uuid.Validate(updateProductDTO.CategoryId) != nil &&
		updateProductDTO.Name == "" &&
		updateProductDTO.Value == 0
	sameData := existingProduct.CategoryId == updateProductDTO.CategoryId &&
		existingProduct.Name == updateProductDTO.Name &&
		existingProduct.Value == updateProductDTO.Value &&
		existingProduct.Description == updateProductDTO.Description
	if emptyData || sameData {
		return entity.Product{}, entity.ErrNoUpdatedData
	}
	if uuid.Validate(updateProductDTO.CategoryId) != nil {
		updateProductDTO.CategoryId = existingProduct.CategoryId
	}
	category, err := s.categoryGateway.GetCategoryById(ctx, updateProductDTO.CategoryId)
	if err != nil {
		if err == entity.ErrDataNotFound {
			return entity.Product{}, err
		}
		return entity.Product{}, fmt.Errorf("cannot update product for this category - %s", err.Error())
	}
	newUpdateProduct := entity.Product{
		Name:        updateProductDTO.Name,
		Description: updateProductDTO.Description,
		Image:       updateProductDTO.Image,
		Value:       updateProductDTO.Value,
		CategoryId:  updateProductDTO.CategoryId,
		Category:    category,
	}
	_, err = s.productGateway.UpdateProduct(ctx, newUpdateProduct)
	if err != nil {
		if err == entity.ErrConflictingData {
			return entity.Product{}, err
		}
		return entity.Product{}, entity.ErrInternal
	}
	return newUpdateProduct, nil
}
