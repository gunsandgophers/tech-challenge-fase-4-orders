package orders

import (
	"errors"
	"tech-challenge-fase-1/internal/core/entities"
	coreerrors "tech-challenge-fase-1/internal/core/errors"
	"tech-challenge-fase-1/internal/core/repositories"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPaymentOrderUseCase(t *testing.T) {

	orderId := uuid.NewString()

	order := entities.RestoreOrder(
		orderId, nil, []*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_AWAITING_PAYMENT,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository := &repositories.MockOrderRepositoryInterface{}

	usecase := NewPaymentOrderUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(order, nil).Once()
	orderRepository.On("Update", mock.Anything).Return(nil).Once()

	err := usecase.Execute(orderId, entities.ORDER_PAYMENT_PAID.String())

	assert.Nil(t, err)
}

func TestPaymentOrderUseCaseWithCheckPaymentErr(t *testing.T) {

	orderId := uuid.NewString()

	orderRepository := &repositories.MockOrderRepositoryInterface{}

	usecase := NewPaymentOrderUseCase(orderRepository)

	err := usecase.Execute(orderId, entities.ORDER_PAYMENT_AWAITING_PAYMENT.String())

	assert.Error(t, err, coreerrors.ErrInvalidPaymentStatus)
}

func TestPaymentOrderUseCaseWithFindOrderByIDErr(t *testing.T) {

	orderId := uuid.NewString()

	orderRepository := &repositories.MockOrderRepositoryInterface{}

	usecase := NewPaymentOrderUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(nil, errors.New("error")).Once()

	err := usecase.Execute(orderId, entities.ORDER_PAYMENT_PAID.String())

	assert.EqualError(t, err, "error")
}

func TestPaymentOrderUseCaseWithPaymentStatusErr(t *testing.T) {

	orderId := uuid.NewString()

	order := entities.RestoreOrder(
		orderId, nil, []*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_PAID,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository := &repositories.MockOrderRepositoryInterface{}

	usecase := NewPaymentOrderUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(order, nil).Once()
	orderRepository.On("Update", mock.Anything).Return(nil).Once()

	err := usecase.Execute(orderId, entities.ORDER_PAYMENT_PAID.String())

	assert.Error(t, err, coreerrors.ErrOrderNotAwaitingPayment)
}

func TestPaymentOrderUseCaseWithPreparationStatusErr(t *testing.T) {

	orderId := uuid.NewString()

	order := entities.RestoreOrder(
		orderId, nil, []*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_AWAITING_PAYMENT,
		entities.ORDER_PREPARATION_CANCELED,
	)

	orderRepository := &repositories.MockOrderRepositoryInterface{}

	usecase := NewPaymentOrderUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(order, nil).Once()
	orderRepository.On("Update", mock.Anything).Return(nil).Once()

	err := usecase.Execute(orderId, entities.ORDER_PAYMENT_PAID.String())

	assert.Error(t, err, coreerrors.ErrOrderNotAwaitingPreparation)
}

func TestPaymentOrderUseCaseWithPaymentRejected(t *testing.T) {

	orderId := uuid.NewString()

	order := entities.RestoreOrder(
		orderId, nil, []*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_AWAITING_PAYMENT,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository := &repositories.MockOrderRepositoryInterface{}

	usecase := NewPaymentOrderUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(order, nil).Once()
	orderRepository.On("Update", mock.Anything).Return(nil).Once()

	err := usecase.Execute(orderId, entities.ORDER_PAYMENT_REJECTED.String())

	assert.Nil(t, err)
}
