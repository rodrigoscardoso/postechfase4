package handler

import (
	"post-tech-challenge-10soat/internal/controllers"
	om "post-tech-challenge-10soat/internal/delivery/http/mapper"
	dto "post-tech-challenge-10soat/internal/dto/order"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderController controllers.OrderController
}

func NewOrderHandler(orderController controllers.OrderController) OrderHandler {
	return OrderHandler{
		orderController,
	}
}

type orderProductRequest struct {
	ProductID   string `json:"product_id" binding:"required,min=1" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Quantity    int    `json:"quantity" binding:"required,number" example:"1"`
	Observation string `json:"observation" binding:"omitempty" example:"Lanche com batata"`
}

type createOrderRequest struct {
	ClientId string                `json:"client_id" binding:"omitempty" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Products []orderProductRequest `json:"products" binding:"required"`
}

// CreateOrder godoc
//
//	@Summary		Criar um novo pedido (checkout)
//	@Description	Cria um novo pedido com o pagamento pendente
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			createOrderRequest	body		createOrderRequest	true	"Criar ordem body"
//	@Success		200					{object}	om.OrderResponse		"Ordem criada"
//	@Failure		400					{object}	ErrorResponse		"Erro de validação"
//	@Failure		500					{object}	ErrorResponse		"Erro interno"
//	@Router			/orders [post]
//	@Security		BearerAuth
func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var request createOrderRequest
	var products []dto.CreateOrderProduct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}
	for _, product := range request.Products {
		products = append(products, dto.CreateOrderProduct{
			ProductId:   product.ProductID,
			Quantity:    product.Quantity,
			Observation: product.Observation,
		})
	}
	oderInfo := dto.CreateOrderDTO{
		ClientId: request.ClientId,
		Products: products,
	}
	o, err := h.orderController.CreateOrder(ctx, oderInfo)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := om.NewOrderResponse(o)
	handleSuccess(ctx, response)
}

type listOrdersRequest struct {
	Limit uint64 `form:"limit" binding:"required,min=1" example:5""`
}

// ListOrders godoc
//
//	@Summary		Lista os pedidos
//	@Description	Lista os pedidos ordenados pelor seu status na seguinte ordem Pronto > Em preparação > Recebido
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int			true	"Limite de pedidos"
//	@Success		200			{object}	om.ListOrdersResponse			"Pedidos listados"
//	@Failure		400			{object}	ErrorResponse	"Erro de validação"
//	@Failure		500			{object}	ErrorResponse	"Erro interno"
//	@Router			/orders [get]
func (h *OrderHandler) ListOrders(ctx *gin.Context) {
	var request listOrdersRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		validationError(ctx, err)
		return
	}
	listOrders, err := h.orderController.ListOrders(ctx, request.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := om.NewListOrdersResponse(listOrders)
	handleSuccess(ctx, response)
}

type getOrderPaymentStatusRequest struct {
	Id string `uri:"id" binding:"required,min=1" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
}

// GetOrderPaymentStatus godoc
//
//	    @Summary     Consultar status de pagamento
//	    @Description Consultar o status de pagamento de um pedido
//	    @Tags        Orders
//	    @Accept      json
//	    @Produce		json
//		@Param	    id	path		string				true	"ID"
//	    @Success		200	{object}    om.OrderPaymentStatusResponse	"Status do pagamento"
//	    @Failure		400	{object}    ErrorResponse	"Erro de validação"
//		@Failure		404	{object}	ErrorResponse   "Pedido não encontrado"
//	    @Router		/orders/{id}/payment-status [get]
func (h *OrderHandler) GetOrderPaymentStatus(ctx *gin.Context) {
	var request getOrderPaymentStatusRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		validationError(ctx, err)
		return
	}
	orderPaymentStatus, err := h.orderController.GetOrderPaymentStatus(ctx, request.Id)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := om.NewOrderPaymentStatusResponse(orderPaymentStatus)
	handleSuccess(ctx, response)
}

type updateOrderStatusRequest struct {
	Id string `uri:"id" binding:"required,min=1" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
}

type updateOrderStatusQuery struct {
	Status string `form:"status" binding:"required,oneof=preparing ready completed" example:"preparing"`
}

// UpdateOrderStatus godoc
//
//	    @Summary     Atualizar status do pedido
//	    @Description Atualizar status do pedido
//	    @Tags        Orders
//	    @Accept      json
//	    @Produce		json
//		@Param	    id	path		string				true	"ID"
//	    @Param			status	query		string	true	"Status do pedido" Enums(preparing, ready, completed)
//	    @Success		200	{object}    om.UpdateOrderStatusResponse	"Status do pagamento"
//	    @Failure		400	{object}    ErrorResponse	"Erro de validação"
//		@Failure		404	{object}	ErrorResponse   "Pedido não encontrado"
//	    @Router		/orders/{id}/status [patch]
func (h *OrderHandler) UpdateOrderStatus(ctx *gin.Context) {
	var request updateOrderStatusRequest
	var query updateOrderStatusQuery
	if err := ctx.ShouldBindUri(&request); err != nil {
		validationError(ctx, err)
		return
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		validationError(ctx, err)
		return
	}
	orderPaymentStatus, err := h.orderController.UpdateOrderStatus(ctx, request.Id, query.Status)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := om.NewOrderUpdateStatusResponse(orderPaymentStatus)
	handleSuccess(ctx, response)
}
