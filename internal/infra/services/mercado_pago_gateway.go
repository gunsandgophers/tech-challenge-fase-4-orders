package services

import (
	"tech-challenge-fase-1/internal/infra/events"
	"tech-challenge-fase-1/internal/core/dtos"
	// "time"
)

type MercadoPagoGateway struct {
	eventManager *events.EventManager
}

func NewMercadoPagoGateway(eventManager *events.EventManager) *MercadoPagoGateway {
	return &MercadoPagoGateway{
		eventManager: eventManager,
	}
}

func (m *MercadoPagoGateway) Execute(order *dtos.OrderDTO, method dtos.MethodType) (*dtos.CheckoutDTO, error) {

	link := "https://www.pngall.com/wp-content/uploads/2/QR-Code-PNG-Images.png"
	total := order.Total

	return &dtos.CheckoutDTO{
		OrderId: order.Id,
		PaymentLink: &link,
		Method:      &method,
		Amount:      &total,
	}, nil

}
