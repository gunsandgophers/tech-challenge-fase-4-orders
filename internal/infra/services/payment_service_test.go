package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePayment(t *testing.T) {

	client := mocks.NewMockHTTPClientInterface(t)

	service := NewPaymentService(client)

	orderID := uuid.NewString()
	amount := 13.37
	paymentLink := "url"
	method := dtos.PIX

	body := dtos.CheckoutDTO{OrderId: orderID,
		PaymentLink: &paymentLink, Method: &method, Amount: &amount,
	}

	raw, _ := json.Marshal(body)

	HTTPResponse := &http.Response{Body: io.NopCloser(bytes.NewReader(raw))}

	client.On("Do", mock.Anything).Return(HTTPResponse, nil)

	response, err := service.CreatePayment(orderID, amount)

	assert.Nil(t, err)
	assert.Equal(t, response.OrderId, orderID)
}

func TestCreatePaymentWithError(t *testing.T) {

	client := mocks.NewMockHTTPClientInterface(t)

	service := NewPaymentService(client)

	orderID := uuid.NewString()
	amount := 13.37

	client.On("Do", mock.Anything).Return(nil, errors.New("error"))

	response, err := service.CreatePayment(orderID, amount)

	assert.Error(t, err)
	assert.Nil(t, response)
}
