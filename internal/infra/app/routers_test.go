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
	"tech-challenge-fase-1/internal/core/queries"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
	httpserver "tech-challenge-fase-1/internal/infra/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	customerService := services.NewMockCustomerService(t)
	productService := services.NewMockProductServiceInterface(t)
	paymentGateway := services.NewMockPaymentGatewayInterface(t)
	orderRepository := repositories.NewMockOrderRepositoryInterface(t)
	orderDisplayListQuery := queries.NewMockOrderDisplayListQueryInterface(t)

	app := NewAPIAppTest(
		customerService, productService,
		paymentGateway, orderRepository,
		orderDisplayListQuery,
	)

	w := httptest.NewRecorder()

	msg, _ := json.Marshal(httpserver.Payload{"msg": "Hello World! :)"})

	req, _ := http.NewRequest("GET", "/api/v1/", strings.NewReader(""))
	app.httpServer.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(msg), w.Body.String())
}

func TestGetStatusPayment(t *testing.T) {
	customerService := services.NewMockCustomerService(t)
	productService := services.NewMockProductServiceInterface(t)
	paymentGateway := services.NewMockPaymentGatewayInterface(t)
	orderRepository := repositories.NewMockOrderRepositoryInterface(t)
	orderDisplayListQuery := queries.NewMockOrderDisplayListQueryInterface(t)

	order := entities.RestoreOrder(
		uuid.NewString(), nil,
		[]*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_PAID,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository.On("FindOrderByID", order.GetId()).Return(order, nil).Once()

	app := NewAPIAppTest(
		customerService, productService,
		paymentGateway, orderRepository,
		orderDisplayListQuery,
	)

	w := httptest.NewRecorder()

	msg, _ := json.Marshal(httpserver.Payload{
		"data": httpserver.Payload{
			"order_id":       order.GetId(),
			"payment_status": order.GetPaymentStatus(),
		},
		"message": "operation: get-payment-status-order successfull",
	})

	req, _ := http.NewRequest(
		"GET",
		fmt.Sprint("/api/v1/order/", order.GetId(), "/payment-status"),
		strings.NewReader(""),
	)
	app.httpServer.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(msg), w.Body.String())
}

func TestGetOrderDisplay(t *testing.T) {
	customerService := services.NewMockCustomerService(t)
	productService := services.NewMockProductServiceInterface(t)
	paymentGateway := services.NewMockPaymentGatewayInterface(t)
	orderRepository := repositories.NewMockOrderRepositoryInterface(t)
	orderDisplayListQuery := queries.NewMockOrderDisplayListQueryInterface(t)

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

	app := NewAPIAppTest(
		customerService, productService,
		paymentGateway, orderRepository,
		orderDisplayListQuery,
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

func TestUpdatePayment(t *testing.T) {
	customerService := services.NewMockCustomerService(t)
	productService := services.NewMockProductServiceInterface(t)
	paymentGateway := services.NewMockPaymentGatewayInterface(t)
	orderRepository := repositories.NewMockOrderRepositoryInterface(t)
	orderDisplayListQuery := queries.NewMockOrderDisplayListQueryInterface(t)

	order := entities.RestoreOrder(
		uuid.NewString(), nil,
		[]*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_AWAITING_PAYMENT,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository.On("FindOrderByID", order.GetId()).Return(order, nil).Once()
	orderRepository.On("Update", order).Return(nil).Once()

	app := NewAPIAppTest(
		customerService, productService,
		paymentGateway, orderRepository,
		orderDisplayListQuery,
	)

	w := httptest.NewRecorder()

	request, _ := json.Marshal(httpserver.Payload{
		"order_id":       order.GetId(),
		"payment_status": entities.ORDER_PAYMENT_PAID,
	})

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/order/payment",
		bytes.NewReader(request),
	)
	app.httpServer.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

func TestUpdatePreparationStatus(t *testing.T) {
	customerService := services.NewMockCustomerService(t)
	productService := services.NewMockProductServiceInterface(t)
	paymentGateway := services.NewMockPaymentGatewayInterface(t)
	orderRepository := repositories.NewMockOrderRepositoryInterface(t)
	orderDisplayListQuery := queries.NewMockOrderDisplayListQueryInterface(t)

	order := entities.RestoreOrder(
		uuid.NewString(), nil,
		[]*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_PAID,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository.On("FindOrderByID", order.GetId()).Return(order, nil).Once()
	orderRepository.On("Update", order).Return(nil).Once()

	app := NewAPIAppTest(
		customerService, productService,
		paymentGateway, orderRepository,
		orderDisplayListQuery,
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