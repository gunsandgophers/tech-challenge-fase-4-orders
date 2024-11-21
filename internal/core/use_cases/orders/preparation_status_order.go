package orders

import (
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/repositories"
)

type PreparationStatusUpdateUseCase struct {
	orderRepository     repositories.OrderRepositoryInterface
}

func NewPreparationStatusUpdateUseCase(
	orderRepository repositories.OrderRepositoryInterface,
) *PreparationStatusUpdateUseCase {
	return &PreparationStatusUpdateUseCase{
		orderRepository:     orderRepository,
	}
}

func (uc *PreparationStatusUpdateUseCase) Execute(
	orderId string,
	preparationStatusString string,
) error {
	order, err := uc.orderRepository.FindOrderByID(orderId)
	if err != nil {
		return err
	}
	preparationStatus := entities.OrderPreparationStatus(preparationStatusString)
	if err := order.SetPreparationStatus(preparationStatus); err != nil {
		return err
	}
	uc.orderRepository.Update(order)
	return nil
}
