package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	entity "post-tech-challenge-10soat/internal/entities"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestHandleError tests the handleError function with different error types
func TestHandleError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedStatus int
		expectedMsg    string
	}{
		{
			name:           "internal error",
			err:            entity.ErrInternal,
			expectedStatus: http.StatusInternalServerError,
			expectedMsg:    entity.ErrInternal.Error(),
		},
		{
			name:           "not found error",
			err:            entity.ErrDataNotFound,
			expectedStatus: http.StatusNotFound,
			expectedMsg:    entity.ErrDataNotFound.Error(),
		},
		{
			name:           "conflict error",
			err:            entity.ErrConflictingData,
			expectedStatus: http.StatusConflict,
			expectedMsg:    entity.ErrConflictingData.Error(),
		},
		{
			name:           "forbidden error",
			err:            entity.ErrForbidden,
			expectedStatus: http.StatusForbidden,
			expectedMsg:    entity.ErrForbidden.Error(),
		},
		{
			name:           "unknown error",
			err:            errors.New("some unknown error"),
			expectedStatus: http.StatusInternalServerError,
			expectedMsg:    "some unknown error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			// Execute
			handleError(ctx, tt.err)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			var response ErrorResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.False(t, response.Success)
			assert.Contains(t, response.Messages, tt.expectedMsg)
		})
	}
}

// TestHandleSuccess tests the handleSuccess function
func TestHandleSuccess(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	testData := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   "123",
		Name: "Test",
	}

	// Execute
	handleSuccess(ctx, testData)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	assert.Equal(t, "Success", response.Message)
	assert.NotNil(t, response.Data)
}

// TestNewResponse tests the newResponse function
func TestNewResponse(t *testing.T) {
	testData := struct{ Name string }{Name: "Test"}
	rsp := newResponse(true, "Test message", testData)

	assert.True(t, rsp.Success)
	assert.Equal(t, "Test message", rsp.Message)
	assert.Equal(t, testData, rsp.Data)
}

// TestNewErrorResponse tests the newErrorResponse function
func TestNewErrorResponse(t *testing.T) {
	errMsgs := []string{"error 1", "error 2"}
	errRsp := newErrorResponse(errMsgs)

	assert.False(t, errRsp.Success)
	assert.Equal(t, errMsgs, errRsp.Messages)
}
