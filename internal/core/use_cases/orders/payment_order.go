package orders

import (
	"strings"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/errors"
	"tech-challenge-fase-1/internal/core/repositories"
)

type PaymentOrderUseCase struct {
	orderRepository repositories.OrderRepositoryInterface
}

func NewPaymentOrderUseCase(
	orderRepository repositories.OrderRepositoryInterface,
) *PaymentOrderUseCase {
	return &PaymentOrderUseCase{
		orderRepository: orderRepository,
	}
}

func (uc *PaymentOrderUseCase) checkValidPaymentStatus(
	paymentStatusString string,
) (entities.OrderPaymentStatus, error) {
	paymentStatus := entities.OrderPaymentStatus(strings.ToUpper(paymentStatusString))
	if paymentStatus != entities.ORDER_PAYMENT_PAID &&
		paymentStatus != entities.ORDER_PAYMENT_REJECTED {
		return paymentStatus, errors.ErrInvalidPaymentStatus
	}

	return paymentStatus, nil
}

func (uc *PaymentOrderUseCase) checkValidOrder(orderId string) (*entities.Order, error) {
	order, err := uc.orderRepository.FindOrderByID(orderId)
	if err != nil {
		return nil, err
	}
	if order.GetPaymentStatus() != entities.ORDER_PAYMENT_AWAITING_PAYMENT {
		return nil, errors.ErrOrderNotAwaitingPayment
	}
	if order.GetPreparationStatus() != entities.ORDER_PREPARATION_AWAITING {
		return nil, errors.ErrOrderNotAwaitingPreparation
	}
	return order, nil
}

func (uc *PaymentOrderUseCase) processOrder(
	order *entities.Order,
	paymentStatus entities.OrderPaymentStatus,
) error {
	if paymentStatus == entities.ORDER_PAYMENT_PAID {
		order.PaymentReceived()
	} else {
		order.PaymentRejected()
	}
	return uc.orderRepository.Update(order)
}

func (uc *PaymentOrderUseCase) Execute(
	orderId string,
	paymentStatusString string,
) error {
	paymentStatus, err := uc.checkValidPaymentStatus(paymentStatusString)
	if err != nil {
		return err
	}
	order, err := uc.checkValidOrder(orderId)
	if err != nil {
		return err
	}
	return uc.processOrder(order, paymentStatus)
}
