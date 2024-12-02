package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
	httpserver "tech-challenge-fase-1/internal/infra/http"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderDisplay(t *testing.T) {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	customerService := mocks.NewMockCustomerService(t)
	productService := mocks.NewMockProductServiceInterface(t)
	paymentService := mocks.NewMockPaymentServiceInterface(t)
	orderRepository := mocks.NewMockOrderRepositoryInterface(t)
	orderDisplayListQuery := mocks.NewMockOrderDisplayListQueryInterface(t)

	orderDisplay := []*dtos.OrderDisplayDTO{
		{
			Id:                uuid.NewString(),
			CustomerId:        nil,
			PreparationStatus: string(entities.ORDER_PREPARATION_AWAITING),
			CreatedAt:         time.Now(),
			Items: []*dtos.OrderItemDisplayDTO{
				{
					Quantity:    1,
					ProductName: "Um produto ai",
				},
			},
		},
	}

	orderDisplayListQuery.On("Execute").Return(orderDisplay, nil).Once()

	app := NewAPIApp(
		httpServer,
		orderRepository,
		orderDisplayListQuery,
		customerService,
		productService,
		paymentService,
	)

	w := httptest.NewRecorder()

	msg := httpserver.Payload{
		"data": httpserver.Payload{
			"orders": []httpserver.Payload{
				{
					"order_id": orderDisplay[0].Id,
					"items": []httpserver.Payload{
						{"quantity": 1, "product_name": "Um produto ai"},
					},
					"preparation_status": "AWAITING",
					"createdAt":          orderDisplay[0].CreatedAt,
				},
			},
		},
		"message": "operation: order-display-list successfull",
	}

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/order/display",
		strings.NewReader(""),
	)
	app.httpServer.ServeHTTP(w, req)

	var response httpserver.Payload

	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, msg["message"], response["message"])
}

func TestUpdatePreparationStatus(t *testing.T) {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	customerService := mocks.NewMockCustomerService(t)
	productService := mocks.NewMockProductServiceInterface(t)
	paymentService := &mocks.MockPaymentServiceInterface{}
	orderRepository := mocks.NewMockOrderRepositoryInterface(t)
	orderDisplayListQuery := mocks.NewMockOrderDisplayListQueryInterface(t)

	order := entities.RestoreOrder(
		uuid.NewString(), nil,
		[]*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_PAID,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository.On("FindOrderByID", order.GetId()).Return(order, nil).Once()
	orderRepository.On("Update", order).Return(nil).Once()

	app := NewAPIApp(
		httpServer,
		orderRepository,
		orderDisplayListQuery,
		customerService,
		productService,
		paymentService,
	)

	w := httptest.NewRecorder()

	request, _ := json.Marshal(
		httpserver.Payload{
			"preparation_status": entities.ORDER_PREPARATION_IN_PREPARARION,
		})

	req, _ := http.NewRequest(
		"PUT",
		fmt.Sprint("/api/v1/order/", order.GetId(), "/preparation-status"),
		bytes.NewReader(request),
	)
	app.httpServer.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
