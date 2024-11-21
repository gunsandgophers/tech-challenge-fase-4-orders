package services

import (
	"tech-challenge-fase-1/internal/core/dtos"
)

type PaymentGatewayInterface interface {
	Execute(order *dtos.OrderDTO, method dtos.MethodType) (*dtos.CheckoutDTO, error)
}
