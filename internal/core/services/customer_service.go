package services

import (
	"tech-challenge-fase-1/internal/core/entities"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
)

type CustomerServiceInterface interface {
	GetCustomerById(id string) (*entities.Customer, error)
	GetCustomerByCPF(cpf *valueobjects.CPF) (*entities.Customer, error)
}
