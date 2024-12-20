package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/infra/config"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type PaymentService struct {
	client httpserver.HTTPClientInterface
}

func NewPaymentService(client httpserver.HTTPClientInterface) *PaymentService {
	return &PaymentService{
		client: client,
	}
}

type CreatePaymentRequest struct {
	Amount float64 `json:"amount"`
}

func (p *PaymentService) CreatePayment(orderID string, amount float64) (*dtos.CheckoutDTO, error) {
	payload, _ := json.Marshal(CreatePaymentRequest{Amount: amount})
	req, err := http.NewRequest("POST", "http://"+config.SERVICE_PAYMENT_URL+"/api/v1/payment/"+orderID, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var checkout dtos.CheckoutDTOResponse
	err = json.Unmarshal(bodyBytes, &checkout)
	if err != nil {
		return nil, err
	}

	return &checkout.Data, nil
}
