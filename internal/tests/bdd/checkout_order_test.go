package app

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/queries"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"tech-challenge-fase-1/internal/infra/app"
	"tech-challenge-fase-1/internal/tests/fixtures"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type appCtxKey struct{}
type depsCtxKey struct{}

type sandwichCtxKey struct{}
type drinkCtxKey struct{}
type sidedishesCtxKey struct{}
type dessetsCtxKey struct{}

type responseCtxKey struct{}

func thereAreProduct(ctx context.Context, category string) (context.Context, error) {

	productCategory := entities.ProductCategory(category)
	product := entities.RestoreProduct(
		uuid.NewString(), "Um produto ai",
		productCategory,
		13.37, "Uma descricao ai", "")

	deps := ctx.Value(depsCtxKey{}).(*DependenciesCheckoutOrder)
	deps.productService.(*mocks.MockProductServiceInterface).
		On("FindProductByID", product.GetId()).Return(product, nil).Once()

	if productCategory == entities.PRODUCT_CATEGORY_SANDWICH {
		return context.WithValue(ctx, sandwichCtxKey{}, product), nil
	} else if productCategory == entities.PRODUCT_CATEGORY_DRINKS {
		return context.WithValue(ctx, drinkCtxKey{}, product), nil
	} else if productCategory == entities.PRODUCT_CATEGORY_SIDEDISHES {
		return context.WithValue(ctx, sidedishesCtxKey{}, product), nil
	}

	return context.WithValue(ctx, dessetsCtxKey{}, product), nil
}

func iOpenOrder(ctx context.Context) (context.Context, error) {

	sandwich, _ := ctx.Value(sandwichCtxKey{}).(*entities.Product)
	drink, _ := ctx.Value(drinkCtxKey{}).(*entities.Product)
	sidedishes, _ := ctx.Value(sidedishesCtxKey{}).(*entities.Product)
	dessets, _ := ctx.Value(dessetsCtxKey{}).(*entities.Product)

	server := ctx.Value(appCtxKey{}).(*app.APIApp)
	deps := ctx.Value(depsCtxKey{}).(*DependenciesCheckoutOrder)

	paymentLink := ""
	pix := dtos.PIX
	amount := sandwich.GetPrice() + drink.GetPrice() +
		sidedishes.GetPrice() + dessets.GetPrice()

	deps.paymentService.(*mocks.MockPaymentServiceInterface).
		On("CreatePayment", mock.Anything, mock.Anything).
		Return(&dtos.CheckoutDTO{
			OrderId:     uuid.NewString(),
			PaymentLink: &paymentLink,
			Method:      &pix,
			Amount:      &amount,
		}, nil).Once()

	deps.orderRepository.(*mocks.MockOrderRepositoryInterface).On("Insert", mock.Anything).
		Return(nil).Once()

	w := httptest.NewRecorder()

	m := make(map[string]interface{})

	m["products_ids"] = []string{
		sandwich.GetId(), drink.GetId(),
		sidedishes.GetId(), dessets.GetId(),
	}

	body, _ := json.Marshal(m)

	req, _ := http.NewRequest("POST", "/api/v1/order/checkout", bytes.NewReader(body))

	server.HTTPServer().ServeHTTP(w, req)

	_ = context.WithValue(ctx, sandwichCtxKey{}, sandwich)
	_ = context.WithValue(ctx, drinkCtxKey{}, drink)
	_ = context.WithValue(ctx, sidedishesCtxKey{}, sidedishes)
	_ = context.WithValue(ctx, dessetsCtxKey{}, dessets)
	return context.WithValue(ctx, responseCtxKey{}, w), nil
}

func thereShouldOpenedOrder(ctx context.Context, items int) error {

	response, _ := ctx.Value(responseCtxKey{}).(*httptest.ResponseRecorder)

	status := response.Result().StatusCode

	if status != 201 {
		return errors.New("invalid status")
	}

	m := make(map[string]interface{})

	err := json.Unmarshal(response.Body.Bytes(), &m)

	if err != nil {
		return nil
	}

	if m["message"] != "operation: checkout-order successfull" {
		return errors.New("error on open order")
	}

	return nil
}

type DependenciesCheckoutOrder struct {
	customerService       services.CustomerServiceInterface
	productService        services.ProductServiceInterface
	paymentService        services.PaymentServiceInterface
	orderRepository       repositories.OrderRepositoryInterface
	orderDisplayListQuery queries.OrderDisplayListQueryInterface
}

func TestFeatures(t *testing.T) {
	customerService := mocks.NewMockCustomerService(t)
	productService := mocks.NewMockProductServiceInterface(t)
	paymentService := mocks.NewMockPaymentServiceInterface(t)
	orderRepository := mocks.NewMockOrderRepositoryInterface(t)
	orderDisplayListQuery := mocks.NewMockOrderDisplayListQueryInterface(t)
	dependencies := &DependenciesCheckoutOrder{
		customerService:       customerService,
		productService:        productService,
		paymentService:        paymentService,
		orderRepository:       orderRepository,
		orderDisplayListQuery: orderDisplayListQuery,
	}
	server := fixtures.NewAPIAppIntegrationTest(
		orderRepository,
		orderDisplayListQuery,
		customerService,
		productService,
		paymentService,
	)

	ctx := context.WithValue(context.Background(), appCtxKey{}, server)
	ctx = context.WithValue(ctx, depsCtxKey{}, dependencies)

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:         "pretty",
			Paths:          []string{"features"},
			DefaultContext: ctx,
			TestingT:       t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Given(`^there are a (\w+)$`, thereAreProduct)
	sc.Given(`^there are a (\w+)$`, thereAreProduct)
	sc.Given(`^there are a (\w+)$`, thereAreProduct)
	sc.Given(`^there are a (\w+)$`, thereAreProduct)

	sc.When(`^I add all items$`, iOpenOrder)
	sc.Then(`^there should be one order with (\d+) items$`, thereShouldOpenedOrder)
}
