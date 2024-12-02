package orders

import (
	"errors"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckoutOrderUseCase(t *testing.T) {

	productsIds := []string{uuid.NewString()}

	product := entities.RestoreProduct(
		productsIds[0], "Um produto",
		entities.PRODUCT_CATEGORY_DESSERTS,
		1337, "Uma descricao", "",
	)

	checkout := &dtos.CheckoutDTO{}

	orderRepository := &repositories.MockOrderRepositoryInterface{}
	customerService := &services.MockCustomerService{}
	productService := &services.MockProductServiceInterface{}
	paymentService := &services.MockPaymentServiceInterface{}

	productService.On("FindProductByID", productsIds[0]).Return(product, nil).Once()
	paymentService.On("CreatePayment", mock.Anything, mock.Anything).Return(checkout, nil).Once()
	orderRepository.On("Insert", mock.Anything).Return(nil).Once()

	usecase := NewCheckoutOrderUseCase(
		orderRepository, customerService,
		productService, paymentService,
	)

	response, _ := usecase.Execute(nil, productsIds)

	assert.Equal(t, checkout, response)
}

func TestCheckoutOrderUseCaseWithCustomer(t *testing.T) {

	productsIds := []string{uuid.NewString()}
	customerId := uuid.NewString()

	product := entities.RestoreProduct(
		productsIds[0], "Um produto",
		entities.PRODUCT_CATEGORY_DESSERTS,
		1337, "Uma descricao", "",
	)

	customer, _ := entities.RestoreCustomer(
		customerId, "Cliente da Silva",
		"email@email.com.br", "11111111111",
	)

	checkout := &dtos.CheckoutDTO{}

	orderRepository := &repositories.MockOrderRepositoryInterface{}
	customerService := &services.MockCustomerService{}
	productService := &services.MockProductServiceInterface{}
	paymentService := &services.MockPaymentServiceInterface{}

	customerService.On("GetCustomerById", customerId).Return(customer, nil).Once()
	productService.On("FindProductByID", productsIds[0]).Return(product, nil).Once()
	paymentService.On("CreatePayment", mock.Anything, mock.Anything).Return(checkout, nil).Once()
	orderRepository.On("Insert", mock.Anything).Return(nil).Once()

	usecase := NewCheckoutOrderUseCase(
		orderRepository, customerService,
		productService, paymentService,
	)

	response, _ := usecase.Execute(&customerId, productsIds)

	assert.Equal(t, checkout, response)
}

func TestCheckoutOrderUseCaseWithCustomerErr(t *testing.T) {

	productsIds := []string{uuid.NewString()}
	customerId := uuid.NewString()

	orderRepository := &repositories.MockOrderRepositoryInterface{}
	customerService := &services.MockCustomerService{}
	productService := &services.MockProductServiceInterface{}
	paymentService := &services.MockPaymentServiceInterface{}

	customerService.On("GetCustomerById", customerId).Return(nil, errors.New("error")).Once()

	usecase := NewCheckoutOrderUseCase(
		orderRepository, customerService,
		productService, paymentService,
	)

	_, err := usecase.Execute(&customerId, productsIds)

	assert.EqualError(t, err, "error")
}

func TestCheckoutOrderUseCaseWithProductErr(t *testing.T) {

	productsIds := []string{uuid.NewString()}

	orderRepository := &repositories.MockOrderRepositoryInterface{}
	customerService := &services.MockCustomerService{}
	productService := &services.MockProductServiceInterface{}
	paymentService := &services.MockPaymentServiceInterface{}

	productService.On("FindProductByID", productsIds[0]).Return(nil, errors.New("error")).Once()

	usecase := NewCheckoutOrderUseCase(
		orderRepository, customerService,
		productService, paymentService,
	)

	_, err := usecase.Execute(nil, productsIds)

	assert.EqualError(t, err, "error")
}

func TestCheckoutOrderUseCaseWithPaymenteErr(t *testing.T) {

	productsIds := []string{uuid.NewString()}

	product := entities.RestoreProduct(
		productsIds[0], "Um produto",
		entities.PRODUCT_CATEGORY_DESSERTS,
		1337, "Uma descricao", "",
	)

	orderRepository := &repositories.MockOrderRepositoryInterface{}
	customerService := &services.MockCustomerService{}
	productService := &services.MockProductServiceInterface{}
	paymentService := &services.MockPaymentServiceInterface{}

	productService.On("FindProductByID", productsIds[0]).Return(product, nil).Once()
	paymentService.On("CreatePayment", mock.Anything, mock.Anything).Return((*dtos.CheckoutDTO)(nil), errors.New("error")).Once()

	usecase := NewCheckoutOrderUseCase(
		orderRepository, customerService,
		productService, paymentService,
	)

	_, err := usecase.Execute(nil, productsIds)

	assert.EqualError(t, err, "error")
}
