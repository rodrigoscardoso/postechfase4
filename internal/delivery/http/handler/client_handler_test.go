package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"post-tech-challenge-10soat/internal/controllers"
	dto "post-tech-challenge-10soat/internal/dto/client"
	entity "post-tech-challenge-10soat/internal/entities"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClientController is a mock of ClientController interface
type MockClientController struct {
	mock.Mock
}

var _ controllers.ClientController = (*MockClientController)(nil)

func (m *MockClientController) CreateClient(ctx context.Context, clientDTO dto.CreateClientDTO) (entity.Client, error) {
	args := m.Called(ctx, clientDTO)
	return args.Get(0).(entity.Client), args.Error(1)
}

func (m *MockClientController) GetClientByCpf(ctx context.Context, cpf string) (entity.Client, error) {
	args := m.Called(ctx, cpf)
	return args.Get(0).(entity.Client), args.Error(1)
}

func (m *MockClientController) GetClientById(ctx context.Context, id string) (entity.Client, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.Client), args.Error(1)
}

func setupTestRouter(handler *ClientHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/clients", handler.CreateClient)
	r.GET("/clients/:cpf", handler.GetClientByCpf)
	return r
}

func TestClientHandler_CreateClient_Success(t *testing.T) {
	// Setup
	mockCtrl := &MockClientController{}
	handler := &ClientHandler{
		clientController: mockCtrl,
	}
	r := setupTestRouter(handler)

	// Mock expectations
	expectedClient := entity.Client{
		Cpf:   "12345678901",
		Name:  "Test User",
		Email: "test@example.com",
	}

	mockCtrl.On("CreateClient", mock.Anything, mock.AnythingOfType("dto.CreateClientDTO")).
		Return(expectedClient, nil)

	// Test request
	reqBody := map[string]string{
		"cpf":   "12345678901",
		"name":  "Test User",
		"email": "test@example.com",
	}
	jsonValue, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/clients", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record response
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"data"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Test User", response.Data.Name)
	assert.Equal(t, "test@example.com", response.Data.Email)

	// Verify mock was called
	mockCtrl.AssertExpectations(t)
}

func TestClientHandler_GetClientByCpf_Success(t *testing.T) {
	// Setup
	mockCtrl := &MockClientController{}
	handler := &ClientHandler{
		clientController: mockCtrl,
	}
	r := setupTestRouter(handler)

	// Mock expectations
	expectedClient := entity.Client{
		Cpf:   "12345678901",
		Name:  "Test User",
		Email: "test@example.com",
	}

	mockCtrl.On("GetClientByCpf", mock.Anything, "12345678901").
		Return(expectedClient, nil)

	// Test request
	req, _ := http.NewRequest("GET", "/clients/12345678901", nil)

	// Record response
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"data"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Test User", response.Data.Name)
	assert.Equal(t, "test@example.com", response.Data.Email)

	// Verify mock was called
	mockCtrl.AssertExpectations(t)
}
