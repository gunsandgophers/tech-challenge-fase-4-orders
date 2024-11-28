package app

import (
	"tech-challenge-fase-1/internal/infra/controllers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	helloController := controllers.NewHelloController()

	orderController := controllers.NewOrderController(
		app.orderRepository,
		app.customerService,
		app.productService,
		app.mercadoPagoGateway,
		app.orderDisplayListQuery,
	)

	baseUrl := "/api/v1"
	app.httpServer.SetBasePath(baseUrl)
	app.httpServer.GET("/", helloController.Index)

	//orders
	app.httpServer.POST("/order/checkout", orderController.Checkout)
	app.httpServer.GET(
		"/order/:order_id/payment-status",
		orderController.GetPaymentStatus,
	)
	app.httpServer.POST("/order/payment", orderController.Payment)
	app.httpServer.GET("/order/display", orderController.OrderDisplayList)
	app.httpServer.PUT(
		"/order/:order_id/preparation-status",
		orderController.OrderPreparationStatusUpdate,
	)

	app.httpServer.SetSwagger("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
