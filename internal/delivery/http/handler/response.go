package handler

import (
	"errors"
	"net/http"
	entity "post-tech-challenge-10soat/internal/entities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var errorStatusMap = map[error]int{
	entity.ErrInternal:        http.StatusInternalServerError,
	entity.ErrDataNotFound:    http.StatusNotFound,
	entity.ErrConflictingData: http.StatusConflict,
	entity.ErrForbidden:       http.StatusForbidden,
}

func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}

func validationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

type ErrorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

func newErrorResponse(errMsgs []string) ErrorResponse {
	return ErrorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}
