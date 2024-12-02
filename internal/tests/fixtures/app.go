package fixtures

import (
	"tech-challenge-fase-1/internal/core/queries"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"tech-challenge-fase-1/internal/infra/app"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

func NewAPIAppIntegrationTest(
	orderRepository repositories.OrderRepositoryInterface,
	orderDisplayListQuery queries.OrderDisplayListQueryInterface,
	customerService services.CustomerServiceInterface,
	productService services.ProductServiceInterface,
	paymentService services.PaymentServiceInterface,
) *app.APIApp {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	return app.NewAPIApp(
		httpServer,
		orderRepository,
		orderDisplayListQuery,
		customerService,
		productService,
		paymentService,
	)
}
