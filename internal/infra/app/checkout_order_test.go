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
	"testing"

	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type appCtxKey struct{}

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

	server := ctx.Value(appCtxKey{}).(*APIApp)
	server.productService.(*services.MockProductServiceInterface).
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

	server := ctx.Value(appCtxKey{}).(*APIApp)

	paymentLink := ""
	pix := dtos.PIX
	amount := sandwich.GetPrice() + drink.GetPrice() +
		sidedishes.GetPrice() + dessets.GetPrice()

	server.mercadoPagoGateway.(*services.MockPaymentGatewayInterface).
		On("Execute", mock.Anything, mock.Anything).
		Return(&dtos.CheckoutDTO{
			OrderId:     uuid.NewString(),
			PaymentLink: &paymentLink,
			Method:      &pix,
			Amount:      &amount,
		}, nil).Once()

	server.orderRepository.(*repositories.MockOrderRepositoryInterface).On("Insert", mock.Anything).
		Return(nil).Once()

	w := httptest.NewRecorder()

	m := make(map[string]interface{})

	m["products_ids"] = []string{
		sandwich.GetId(), drink.GetId(),
		sidedishes.GetId(), dessets.GetId(),
	}

	body, _ := json.Marshal(m)

	req, _ := http.NewRequest("POST", "/api/v1/order/checkout", bytes.NewReader(body))

	server.httpServer.ServeHTTP(w, req)

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

func TestFeatures(t *testing.T) {

	server := NewAPIAppTest(
		services.NewMockCustomerService(t),
		services.NewMockProductServiceInterface(t),
		services.NewMockPaymentGatewayInterface(t),
		repositories.NewMockOrderRepositoryInterface(t),
		queries.NewMockOrderDisplayListQueryInterface(t),
	)

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:         "pretty",
			Paths:          []string{"../../../features"},
			DefaultContext: context.WithValue(context.Background(), appCtxKey{}, server),
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
