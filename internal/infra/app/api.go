package app

import (
	"tech-challenge-fase-1/internal/core/queries"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type APIApp struct {
	httpServer            httpserver.HTTPServer
	orderRepository       repositories.OrderRepositoryInterface
	orderDisplayListQuery queries.OrderDisplayListQueryInterface
	customerService       services.CustomerServiceInterface
	productService        services.ProductServiceInterface
	paymentService        services.PaymentServiceInterface
}

func NewAPIApp(
	httpServer httpserver.HTTPServer,
	orderRepository repositories.OrderRepositoryInterface,
	orderDisplayListQuery queries.OrderDisplayListQueryInterface,
	customerService services.CustomerServiceInterface,
	productService services.ProductServiceInterface,
	paymentService services.PaymentServiceInterface,
) *APIApp {
	app := &APIApp{}
	// // HTTP SERVER
	app.httpServer = httpServer

	// REPOSITORIES AND SERVICES
	app.orderRepository = orderRepository
	app.orderDisplayListQuery = orderDisplayListQuery
	app.customerService = customerService
	app.productService = productService
	app.paymentService = paymentService

	// ROUTES
	app.configRoutes()
	return app
}

func (app *APIApp) configRoutes() {
	registerRouters(app)
}

func (app *APIApp) HTTPServer() httpserver.HTTPServer {
	return app.httpServer
}

func (app *APIApp) Run() {
	app.httpServer.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
