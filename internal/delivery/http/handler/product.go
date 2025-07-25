package handler

import (
	"fmt"
	"post-tech-challenge-10soat/internal/controllers"
	pm "post-tech-challenge-10soat/internal/delivery/http/mapper"
	dto "post-tech-challenge-10soat/internal/dto/product"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productController controllers.ProductController
}

func NewProductHandler(productController controllers.ProductController) ProductHandler {
	return ProductHandler{
		productController,
	}
}

type listProductsRequest struct {
	CategoryID string `form:"category_id" binding:"omitempty,min=1" example:ed6ac028-8016-4cbd-aeee-c3a155cdb2a4""`
}

// ListProducts godoc
//
//	@Summary		Lista os produtos
//	@Description	Lista os produtos podendo buscar por categoria
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			category_id	query		string			false	"Id da categoria"
//	@Success		200			{array}	pm.ProductResponse			"Produtos listados"
//	@Failure		400			{object}	ErrorResponse	"Erro de validação"
//	@Failure		500			{object}	ErrorResponse	"Erro interno"
//	@Router			/products [get]
func (h *ProductHandler) ListProducts(ctx *gin.Context) {
	var request listProductsRequest
	var productsList []pm.ProductResponse
	if err := ctx.ShouldBindQuery(&request); err != nil {
		validationError(ctx, err)
		return
	}
	products, err := h.productController.ListProducts(ctx, request.CategoryID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, p := range products {
		productsList = append(productsList, pm.NewProductResponse(p))
	}
	handleSuccess(ctx, productsList)
}

type createProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"Lanche"`
	Description string  `json:"description" binding:"omitempty" example:"Lanche com batata"`
	Image       string  `json:"image" binding:"omitempty" example:"https://"`
	Value       float64 `json:"value" binding:"required" example:"10.90"`
	CategoryID  string  `json:"category_id" binding:"omitempty,min=1" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
}

// CreateProduct godoc
//
//	@Summary     Registra um novo produto
//	@Description registra um novo produto
//	@Tags        Products
//	@Accept      json
//	@Produce		json
//	@Param	    createProductRequest	body createProductRequest true "Registrar novo produto body"
//	@Success		200	{object} pm.ProductResponse	"Produto registrado"
//	@Failure		400	{object} ErrorResponse	"Erro de validação"
//	@Router		/products [post]
func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var request createProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}
	categoryId, err := uuid.Parse(request.CategoryID)
	if err != nil {
		handleError(ctx, fmt.Errorf("invalid category id"))
		return
	}
	newProduct := dto.CreateProductDTO{
		Name:        request.Name,
		Description: request.Description,
		Image:       request.Image,
		Value:       request.Value,
		CategoryId:  categoryId.String(),
	}
	product, err := h.productController.CreateProduct(ctx, newProduct)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, pm.NewProductResponse(product))
}

type updateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"Lanche"`
	Description string  `json:"description" binding:"omitempty" example:"Lanche com batata"`
	Image       string  `json:"image" binding:"omitempty" example:"https://"`
	Value       float64 `json:"value" binding:"required" example:"10.90"`
	CategoryID  string  `json:"category_id" binding:"omitempty,min=1" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
}

// UpdateProduct godoc
//
//	@Summary     Atualiza um produto
//	@Description Atualiza um produto
//	@Tags        Products
//	@Accept      json
//	@Produce		json
//	@Param			id						path		string					true	"Id do produto"
//	@Param			updateProductRequest	body		updateProductRequest	true	"Atualizar produto body"
//	@Success		200	{object} pm.ProductResponse	"Produto atualizado"
//	@Failure		404						{object}	ErrorResponse			"Produto nao encontrado"
//	@Failure		400	{object} ErrorResponse	"Erro de validação"
//	@Router		/products/{id} [put]
func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var request updateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}
	categoryId, err := uuid.Parse(request.CategoryID)
	if err != nil {
		handleError(ctx, fmt.Errorf("invalid category id"))
		return
	}
	productId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		handleError(ctx, fmt.Errorf("invalid product id"))
		return
	}
	updateProduct := dto.UpdateProductDTO{
		Id:          productId.String(),
		Name:        request.Name,
		Description: request.Description,
		Image:       request.Image,
		Value:       request.Value,
		CategoryId:  categoryId.String(),
	}
	product, err := h.productController.UpdateProduct(ctx, updateProduct)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, pm.NewProductResponse(product))
}

type deleteProductRequest struct {
	Id string `uri:"id" binding:"required,min=1" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
}

// UpdateProduct godoc
//
//	@Summary     Remove um produto
//	@Description Remove um produto por meio de seu identificador
//	@Tags        Products
//	@Accept      json
//	@Produce		json
//	@Param			id						path		string					true	"Id do produto"
//	@Success		200	{object} pm.ProductResponse	"Produto removido"
//	@Failure		404						{object}	ErrorResponse			"Produto nao encontrado"
//	@Failure		400	{object} ErrorResponse	"Erro de validação"
//	@Router		/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	var request deleteProductRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		validationError(ctx, err)
		return
	}
	err := h.productController.DeleteProduct(ctx, request.Id)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, nil)
}
