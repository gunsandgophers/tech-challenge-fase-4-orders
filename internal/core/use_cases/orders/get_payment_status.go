package orders

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/repositories"
)

type GetPaymentStatusUseCase struct {
	orderRepository     repositories.OrderRepositoryInterface
}

func NewGetPaymentStatusUseCase(
	orderRepository repositories.OrderRepositoryInterface,
) *GetPaymentStatusUseCase {
	return &GetPaymentStatusUseCase{
		orderRepository:     orderRepository,
	}
}

func (ps *GetPaymentStatusUseCase) Execute(
	orderId string,
) (*dtos.PaymentStatusDTO, error) {
	order, err := ps.orderRepository.FindOrderByID(orderId)
	if err != nil {
		return nil, err
	}
	return dtos.NewPaymentStatusDTOFromEntity(order), nil
}

