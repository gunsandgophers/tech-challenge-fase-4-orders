package main

import (
	"net/http"
	_ "tech-challenge-fase-1/docs"
	"tech-challenge-fase-1/internal/infra/app"
	"tech-challenge-fase-1/internal/infra/config"
	"tech-challenge-fase-1/internal/infra/database"
	httpserver "tech-challenge-fase-1/internal/infra/http"
	"tech-challenge-fase-1/internal/infra/queries"
	"tech-challenge-fase-1/internal/infra/repositories"
	"tech-challenge-fase-1/internal/infra/services"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	db := database.ConnectDatabase()
	orderRepository := repositories.NewOrderRepositoryMongo(db)
	orderDisplayListQuery := queries.NewOrderDisplayListQueryMongo(db)

	client, err := services.NewCognito(config.AWS_REGION)
	if err != nil {
		panic(err)
	}

	customerService := services.NewAwsCustomerService(client, config.AWS_USER_POOL_ID)

	productService := services.NewProductService(http.DefaultClient)
	paymentService := services.NewPaymentService(http.DefaultClient)

	app := app.NewAPIApp(
		httpServer,
		orderRepository,
		orderDisplayListQuery,
		customerService,
		productService,
		paymentService,
	)
	app.Run()
}
