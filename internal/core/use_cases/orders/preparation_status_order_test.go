package orders

import (
	"errors"
	"tech-challenge-fase-1/internal/core/entities"
	coreerrors "tech-challenge-fase-1/internal/core/errors"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPreparationStatusUpdateUseCase(t *testing.T) {

	orderId := uuid.NewString()

	order := entities.RestoreOrder(
		orderId, nil, []*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_AWAITING_PAYMENT,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository := mocks.NewMockOrderRepositoryInterface(t)

	usecase := NewPreparationStatusUpdateUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(order, nil).Once()

	order.SetPreparationStatus(entities.ORDER_PREPARATION_IN_PREPARARION)
	orderRepository.On("Update", order).Return(nil).Once()

	err := usecase.Execute(orderId, entities.ORDER_PREPARATION_IN_PREPARARION.String())

	assert.Nil(t, err)
}

func TestPreparationStatusUpdateUseCaseWithFindOrderByIDErr(t *testing.T) {

	orderId := uuid.NewString()

	orderRepository := mocks.NewMockOrderRepositoryInterface(t)

	usecase := NewPreparationStatusUpdateUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(nil, errors.New("error")).Once()

	err := usecase.Execute(orderId, entities.ORDER_PREPARATION_IN_PREPARARION.String())

	assert.EqualError(t, err, "error")
}

func TestPreparationStatusUpdateUseCaseWithStatusIncorrect(t *testing.T) {

	orderId := uuid.NewString()

	order := entities.RestoreOrder(
		orderId, nil, []*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_AWAITING_PAYMENT,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository := mocks.NewMockOrderRepositoryInterface(t)

	usecase := NewPreparationStatusUpdateUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(order, nil).Once()

	err := usecase.Execute(orderId, "banana")

	assert.Error(t, err, coreerrors.ErrInvalidPreparationStatus)
}

func TestPreparationStatusUpdateUseCaseWithUpdateErr(t *testing.T) {

	orderId := uuid.NewString()

	order := entities.RestoreOrder(
		orderId, nil, []*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_AWAITING_PAYMENT,
		entities.ORDER_PREPARATION_AWAITING,
	)

	orderRepository := mocks.NewMockOrderRepositoryInterface(t)

	usecase := NewPreparationStatusUpdateUseCase(orderRepository)

	orderRepository.On("FindOrderByID", orderId).Return(order, nil).Once()

	order.SetPreparationStatus(entities.ORDER_PREPARATION_IN_PREPARARION)
	orderRepository.On("Update", order).Return(errors.New("error")).Once()

	err := usecase.Execute(orderId, entities.ORDER_PREPARATION_IN_PREPARARION.String())

	assert.EqualError(t, err, "error")
}
