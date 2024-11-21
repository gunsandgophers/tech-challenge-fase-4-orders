package repositories

import "tech-challenge-fase-1/internal/core/entities"

type OrderRepositoryInterface interface {
	Insert(order *entities.Order) error
	Update(order *entities.Order) error
	FindOrderByID(orderID string) (*entities.Order, error)
}
