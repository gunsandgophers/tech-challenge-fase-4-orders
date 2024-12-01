package fixtures

import (
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"tech-challenge-fase-1/internal/infra/app"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)


func NewAPIAppBDDTest(
	paymentRepository repositories.PaymentRepositoryInterface,
	paymentGateway services.PaymentGatewayInterface,
) *app.APIApp {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	return app.NewAPIApp(httpServer, paymentRepository, paymentGateway)
}


