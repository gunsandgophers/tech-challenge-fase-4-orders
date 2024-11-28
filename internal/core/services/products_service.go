package services

import (
	"tech-challenge-fase-1/internal/core/entities"
)

type ProductServiceInterface interface {
	FindProductByID(id string) (*entities.Product, error)
}
