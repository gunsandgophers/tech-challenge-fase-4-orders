package orders

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/queries"
)

type OrderDisplayListUseCase struct {
	orderDisplayListQuery queries.OrderDisplayListQueryInterface
}

func NewOrderDisplayListUseCase(
	orderDisplayListQuery queries.OrderDisplayListQueryInterface,
) *OrderDisplayListUseCase {
	return &OrderDisplayListUseCase{
		orderDisplayListQuery: orderDisplayListQuery,
	}
}

func (uc *OrderDisplayListUseCase) Execute() ([]*dtos.OrderDisplayDTO, error) {
	return uc.orderDisplayListQuery.Execute()
}
