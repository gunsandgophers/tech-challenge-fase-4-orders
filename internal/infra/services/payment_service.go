package services

import (
	"net/http"
	"tech-challenge-fase-1/internal/core/dtos"
)

type PaymentService struct {
	client *http.Client
}

func NewPaymentService(client *http.Client) *PaymentService {
	return &PaymentService{
		client: client,
	}
}

func (p *PaymentService) CreatePayment(orderID string, amount float64) (*dtos.CheckoutDTO, error) {
	return &dtos.CheckoutDTO{}, nil
}
