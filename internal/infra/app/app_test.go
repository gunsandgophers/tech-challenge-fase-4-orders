package app

import (
	"tech-challenge-fase-1/internal/core/queries"
	"tech-challenge-fase-1/internal/core/repositories"
	coreservices "tech-challenge-fase-1/internal/core/services"
)

func NewAPIAppTest(
	customerService coreservices.CustomerServiceInterface,
	productService coreservices.ProductServiceInterface,
	paymentGateway coreservices.PaymentGatewayInterface,
	orderRepository repositories.OrderRepositoryInterface,
	orderDisplayListQuery queries.OrderDisplayListQueryInterface,

) *APIApp {
	app := &APIApp{}

	app.customerService = customerService
	app.productService = productService
	app.mercadoPagoGateway = paymentGateway
	app.orderRepository = orderRepository
	app.orderDisplayListQuery = orderDisplayListQuery

	app.initGin()
	app.configCors()
	app.configRoutes()

	return app
}
