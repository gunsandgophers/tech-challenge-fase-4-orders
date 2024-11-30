package app

import (
	"net/http"
	corequeries "tech-challenge-fase-1/internal/core/queries"
	corerepositories "tech-challenge-fase-1/internal/core/repositories"
	coreservices "tech-challenge-fase-1/internal/core/services"

	"tech-challenge-fase-1/internal/infra/config"
	"tech-challenge-fase-1/internal/infra/database"

	httpserver "tech-challenge-fase-1/internal/infra/http"
	"tech-challenge-fase-1/internal/infra/queries"
	"tech-challenge-fase-1/internal/infra/repositories"
	"tech-challenge-fase-1/internal/infra/services"

	"github.com/gin-contrib/cors"
)

type APIApp struct {
	httpServer            *httpserver.GinHTTPServerAdapter
	connection            *database.PGXConnectionAdapter
	customerService       coreservices.CustomerServiceInterface
	productService        coreservices.ProductServiceInterface
	orderRepository       corerepositories.OrderRepositoryInterface
	orderDisplayListQuery corequeries.OrderDisplayListQueryInterface
	mercadoPagoGateway    coreservices.PaymentGatewayInterface
}

func NewAPIApp() *APIApp {
	app := &APIApp{}
	app.initGin()
	app.configCors()
	app.initConnectionDB()
	app.initExternalServices()
	app.configRoutes()
	return app
}

func (app *APIApp) initGin() {
	app.httpServer = httpserver.NewGinHTTPServerAdapter()
	app.httpServer.SetTrustedProxies(nil)
}

func (app *APIApp) configCors() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}
	app.httpServer.Engine.Use(cors.New(config))
}

func (app *APIApp) initConnectionDB() {
	app.connection = database.NewPGXConnectionAdapter()

	app.orderRepository = repositories.NewOrderRepositoryDB(app.connection)
	app.orderDisplayListQuery = queries.NewOrderDisplayListQueryDB(app.connection)
}

func (app *APIApp) initExternalServices() {

	app.mercadoPagoGateway = services.NewMercadoPagoGateway()

	var err error
	app.customerService, err = services.NewAwsCustomerService(
		config.AWS_REGION,
		config.AWS_USER_POOL_ID,
	)

	app.productService = services.NewProductService(http.DefaultClient)

	if err != nil {
		panic(err)
	}
}

func (app *APIApp) configRoutes() {
	registerRouters(app)
}

func (app *APIApp) Run() {
	app.httpServer.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (app *APIApp) Shutdown() {
	app.connection.Close()
}
