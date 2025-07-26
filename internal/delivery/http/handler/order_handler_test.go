package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"post-tech-challenge-10soat/internal/controllers"
	dto "post-tech-challenge-10soat/internal/dto/order"
	entity "post-tech-challenge-10soat/internal/entities"
	"post-tech-challenge-10soat/internal/usecases/order"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock use cases for OrderController
type MockCreateOrderUseCase struct {
	mock.Mock
}

func (m *MockCreateOrderUseCase) Execute(ctx context.Context, createOrder dto.CreateOrderDTO) (entity.Order, error) {
	args := m.Called(ctx, createOrder)
	return args.Get(0).(entity.Order), args.Error(1)
}

type MockListOrdersUseCase struct {
	mock.Mock
}

func (m *MockListOrdersUseCase) Execute(ctx context.Context, limit uint64) ([]entity.Order, error) {
	args := m.Called(ctx, limit)
	return args.Get(0).([]entity.Order), args.Error(1)
}

type MockGetOrderPaymentStatusUseCase struct {
	mock.Mock
}

func (m *MockGetOrderPaymentStatusUseCase) Execute(ctx context.Context, id string) (order.OrderPaymentStatus, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(order.OrderPaymentStatus), args.Error(1)
}

type MockUpdateOrderStatusUseCase struct {
	mock.Mock
}

func (m *MockUpdateOrderStatusUseCase) Execute(ctx context.Context, id string, status string) (entity.Order, error) {
	args := m.Called(ctx, id, status)
	return args.Get(0).(entity.Order), args.Error(1)
}

// setupTestController creates a real OrderController with mock use cases
func setupTestController() (*controllers.OrderController, *MockCreateOrderUseCase, *MockListOrdersUseCase, *MockGetOrderPaymentStatusUseCase, *MockUpdateOrderStatusUseCase) {
	mockCreateOrder := &MockCreateOrderUseCase{}
	mockListOrders := &MockListOrdersUseCase{}
	mockGetOrderPaymentStatus := &MockGetOrderPaymentStatusUseCase{}
	mockUpdateOrderStatus := &MockUpdateOrderStatusUseCase{}

	controller := controllers.NewOrderController(
		mockCreateOrder,
		mockListOrders,
		mockGetOrderPaymentStatus,
		mockUpdateOrderStatus,
	)

	return controller, mockCreateOrder, mockListOrders, mockGetOrderPaymentStatus, mockUpdateOrderStatus
}

func setupOrderTestRouter(handler *OrderHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders", handler.ListOrders)
	r.GET("/orders/:id/payment-status", handler.GetOrderPaymentStatus)
	r.PATCH("/orders/:id/status", handler.UpdateOrderStatus)
	return r
}

func TestOrderHandler_CreateOrder_Success(t *testing.T) {
	// Setup
	controller, mockCreateOrder, _, _, _ := setupTestController()
	handler := &OrderHandler{
		orderController: *controller,
	}
	r := setupOrderTestRouter(handler)

	// Mock expectations
	expectedOrder := entity.Order{
		Id:        uuid.NewString(),
		ClientId:  uuid.NewString(),
		Status:    "received",
		CreatedAt: time.Now(),
	}

	mockCreateOrder.On("Execute", mock.Anything, mock.Anything).Return(expectedOrder, nil)

	// Test request
	reqBody := map[string]interface{}{
		"client_id": expectedOrder.ClientId,
		"products": []map[string]interface{}{
			{
				"product_id": uuid.NewString(),
				"quantity":   2,
			},
		},
	}

	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Execute
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Data.ID)

	// Verify mock was called
	mockCreateOrder.AssertExpectations(t)
}

func TestOrderHandler_ListOrders_Success(t *testing.T) {
	// Setup
	controller, _, mockListOrders, _, _ := setupTestController()
	handler := &OrderHandler{
		orderController: *controller,
	}
	r := setupOrderTestRouter(handler)

	// Mock expectations
	expectedOrders := []entity.Order{
		{
			Id:        uuid.NewString(),
			ClientId:  uuid.NewString(),
			Status:    "received",
			CreatedAt: time.Now(),
		},
	}

	mockListOrders.On("Execute", mock.Anything, uint64(10)).Return(expectedOrders, nil)

	// Test request
	req, _ := http.NewRequest("GET", "/orders?limit=10", nil)

	// Execute
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response.Data, 1)

	// Verify mock was called
	mockListOrders.AssertExpectations(t)
}

func TestOrderHandler_GetOrderPaymentStatus_Success(t *testing.T) {
	// Setup
	controller, _, _, mockGetOrderPaymentStatus, _ := setupTestController()
	handler := &OrderHandler{
		orderController: *controller,
	}
	r := setupOrderTestRouter(handler)

	// Mock expectations
	orderID := uuid.NewString()
	expectedStatus := order.OrderPaymentStatus{
		PaymentStatus: order.PaymentPending,
	}

	mockGetOrderPaymentStatus.On("Execute", mock.Anything, orderID).Return(expectedStatus, nil)

	// Test request
	req, _ := http.NewRequest("GET", "/orders/"+orderID+"/payment-status", nil)

	// Execute
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data struct {
			PaymentStatus string `json:"paymentStatus"`
		} `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, string(order.PaymentPending), response.Data.PaymentStatus)

	// Verify mock was called
	mockGetOrderPaymentStatus.AssertExpectations(t)
}

func TestOrderHandler_UpdateOrderStatus_Success(t *testing.T) {
	// Setup
	controller, _, _, _, mockUpdateOrderStatus := setupTestController()
	handler := &OrderHandler{
		orderController: *controller,
	}
	r := setupOrderTestRouter(handler)

	// Mock expectations
	orderID := uuid.NewString()
	status := "preparing"
	expectedOrder := entity.Order{
		Id:        orderID,
		Status:    "preparing",
		UpdatedAt: time.Now(),
	}

	mockUpdateOrderStatus.On("Execute", mock.Anything, orderID, status).Return(expectedOrder, nil)

	// Test request
	req, _ := http.NewRequest("PATCH", "/orders/"+orderID+"/status?status="+status, nil)

	// Execute
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data struct {
			ID     string `json:"id"`
			Status string `json:"status"`
		} `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, orderID, response.Data.ID)
	assert.Equal(t, status, response.Data.Status)

	// Verify mock was called
	mockUpdateOrderStatus.AssertExpectations(t)
}
