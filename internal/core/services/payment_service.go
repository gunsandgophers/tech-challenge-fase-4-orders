package services

import "tech-challenge-fase-1/internal/core/dtos"

type PaymentServiceInterface interface {
	CreatePayment(orderID string, amount float64) (*dtos.CheckoutDTO, error)
}

// OrderId     string
// PaymentLink string
// Method      MethodType
// Amount      float64
