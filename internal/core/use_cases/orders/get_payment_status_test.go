package orders

import (
	"errors"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/repositories"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentStatusUseCase(t *testing.T) {

	mockOrderRepository := &repositories.MockOrderRepositoryInterface{}

	orderId := uuid.NewString()

	order := entities.RestoreOrder(
		orderId, nil, []*valueobjects.OrderItem{},
		entities.ORDER_PAYMENT_PAID,
		entities.ORDER_PREPARATION_AWAITING)

	useCase := NewGetPaymentStatusUseCase(mockOrderRepository)

	mockOrderRepository.On("FindOrderByID", orderId).Return(order, nil).Once()

	paymentStatusDTO, _ := useCase.Execute(orderId)

	assert.Equal(t, paymentStatusDTO.OrderId, orderId)
	assert.Equal(t, paymentStatusDTO.PaymentStatus, entities.ORDER_PAYMENT_PAID.String())
}

func TestGetPaymentStatusUseCaseWithErr(t *testing.T) {

	mockOrderRepository := &repositories.MockOrderRepositoryInterface{}

	orderId := uuid.NewString()

	useCase := NewGetPaymentStatusUseCase(mockOrderRepository)

	mockOrderRepository.On("FindOrderByID", orderId).Return(nil, errors.New("error")).Once()

	_, err := useCase.Execute(orderId)

	assert.EqualError(t, err, "error", "error")
}
