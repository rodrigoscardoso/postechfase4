package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupHealthTestRouter(handler *HealthHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/health", handler.HealthCheck)
	return r
}

func TestHealthHandler_HealthCheck(t *testing.T) {
	// Setup
	handler := NewHealthHandler()
	r := setupHealthTestRouter(&handler)

	// Test request
	req, _ := http.NewRequest("GET", "/health", nil)

	// Execute
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Application is healthy", response["message"])
}
