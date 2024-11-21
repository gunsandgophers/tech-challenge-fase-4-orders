package dtos

import "tech-challenge-fase-1/internal/core/entities"

type OrderDTO struct {
	Id     string         `json:"order_id,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	Items  []*OrderItemDTO `json:"items,omitempty"`
	PaymentStatus string         `json:"payment_status,omitempty"`
	PreparationStatus string         `json:"preparation_status,omitempty"`
	Total  float64        `json:"total,omitempty"`
}

type OrderItemDTO struct {
	Amount      float64 `json:"amount,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	ProductName string  `json:"product_name,omitempty"`
}

func NewOrderDTOFromEntity(order *entities.Order) *OrderDTO {
	orderItems := []*OrderItemDTO{}
	for _, item := range order.GetItems() {
		orderItems = append(orderItems, &OrderItemDTO{
			Amount:      item.GetAmount(),
			Quantity:    item.GetQuantity(),
			ProductName: item.GetProductName(),
		})
	}
	return &OrderDTO{
		Id:     order.GetId(),
		CustomerId: order.GetCustomerId(),
		Items:  orderItems,
		PaymentStatus: order.GetPaymentStatus().String(),
		PreparationStatus: order.GetPreparationStatus().String(),
		Total:  order.GetTotal(),
	}
}
