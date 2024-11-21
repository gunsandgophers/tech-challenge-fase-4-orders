package dtos

import "time"

type OrderDisplayDTO struct {
	Id                string                 `json:"order_id,omitempty"`
	CustomerId        *string                `json:"customer_id,omitempty"`
	Items             []*OrderItemDisplayDTO `json:"items,omitempty"`
	PreparationStatus string                 `json:"preparation_status,omitempty"`
	CreatedAt         time.Time              `json:"createdAt,omitempty"`
}

type OrderItemDisplayDTO struct {
	Quantity    int    `json:"quantity,omitempty"`
	ProductName string `json:"product_name,omitempty"`
}
