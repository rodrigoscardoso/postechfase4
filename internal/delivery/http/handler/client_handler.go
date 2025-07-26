package handler

import (
	"post-tech-challenge-10soat/internal/controllers"
	cm "post-tech-challenge-10soat/internal/delivery/http/mapper"
	dto "post-tech-challenge-10soat/internal/dto/client"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientController controllers.ClientController
}

func NewClientHandler(clientController controllers.ClientController) ClientHandler {
	return ClientHandler{
		clientController: clientController,
	}
}

type createClientRequest struct {
	Cpf   string `json:"cpf" binding:"required" example:"12345678010"`
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required" example:"john-doe@email.com"`
}

// CreateClient godoc
//
//	@Summary     Registra um novo cliente
//	@Description Registra um novo cliente com nome e e-mail
//	@Tags        Clients
//	@Accept      json
//	@Produce		json
//	@Param	    createClientRequest	body createClientRequest true "Registrar novo cliente request"
//	@Success		200	{object} cm.ClientResponse	"Cliente registrado"
//	@Failure		400	{object} ErrorResponse	"Erro de validação"
//	@Router		/clients [post]
func (h *ClientHandler) CreateClient(ctx *gin.Context) {
	var request createClientRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}
	newClient := dto.CreateClientDTO{
		Cpf:   request.Cpf,
		Name:  request.Name,
		Email: request.Email,
	}
	createdClient, err := h.clientController.CreateClient(ctx, newClient)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := cm.NewClientResponse(createdClient)
	handleSuccess(ctx, response)
}

type getClientByCpfRequest struct {
	Cpf string `uri:"cpf" binding:"required,min=1" example:"12345678010"`
}

// GetClientByCpf godoc
//
//	    @Summary     Busca um cliente
//	    @Description buscar um cliente pelo Cpf
//	    @Tags        Clients
//	    @Accept      json
//	    @Produce		json
//		   @Param	    cpf	path		string				true	"CPF"
//	    @Success		200	{object}    cm.ClientResponse	"Cliente"
//	    @Failure		400	{object}    ErrorResponse	"Erro de validação"
//		   @Failure		404	{object}	ErrorResponse   "Cliente nao encontrado"
//	    @Router		/clients/{cpf} [get]
func (h *ClientHandler) GetClientByCpf(ctx *gin.Context) {
	var request getClientByCpfRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		validationError(ctx, err)
		return
	}
	c, err := h.clientController.GetClientByCpf(ctx, request.Cpf)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := cm.NewClientResponse(c)
	handleSuccess(ctx, response)
}
