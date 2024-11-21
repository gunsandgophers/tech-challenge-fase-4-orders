package dtos

import "tech-challenge-fase-1/internal/core/entities"

type PaymentStatusDTO struct {
	OrderId       string `json:"order_id"`
	PaymentStatus string `json:"payment_status"`
}

func NewPaymentStatusDTOFromEntity(order *entities.Order) *PaymentStatusDTO {
	return &PaymentStatusDTO{
		OrderId:       order.GetId(),
		PaymentStatus: order.GetPaymentStatus().String(),
	}
}
