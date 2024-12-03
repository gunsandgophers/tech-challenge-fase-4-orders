package app

import (
	"tech-challenge-fase-1/internal/infra/controllers"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	// helloController := controllers.NewHelloController()

	orderController := controllers.NewOrderController(
		app.orderRepository,
		app.customerService,
		app.productService,
		app.paymentService,
		app.orderDisplayListQuery,
	)

	baseUrl := "/api/v1"
	app.httpServer.(httpserver.HTTPRoutes).SetBasePath(baseUrl)

	//orders
	app.httpServer.(httpserver.HTTPRoutes).POST("/order/checkout", orderController.Checkout)
	app.httpServer.(httpserver.HTTPRoutes).GET("/order/display", orderController.OrderDisplayList)
	app.httpServer.(httpserver.HTTPRoutes).PUT(
		"/order/:order_id/preparation-status",
		orderController.OrderPreparationStatusUpdate,
	)
	app.httpServer.(httpserver.HTTPRoutes).SetSwagger("/swagger/*any")
}
