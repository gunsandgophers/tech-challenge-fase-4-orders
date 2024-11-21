package queries

import (
	"tech-challenge-fase-1/internal/core/dtos"
)

type OrderDisplayListQueryInterface interface {
	Execute() ([]*dtos.OrderDisplayDTO, error)
}
