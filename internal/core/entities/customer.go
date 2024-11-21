package entities

import (
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"

	uuid "github.com/google/uuid"
)

type Customer struct {
	id    string
	name  string
	email *valueobjects.Email
	cpf   *valueobjects.CPF
}

func CreateCustomer(name string, email string, cpf string) (*Customer, error) {
	emailVO, err := valueobjects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	cpfVO, err := valueobjects.NewCPF(cpf)
	if err != nil {
		return nil, err
	}

	return &Customer{
		id:    uuid.NewString(),
		name:  name,
		email: emailVO,
		cpf:   cpfVO,
	}, nil
}

func RestoreCustomer(id string, name string, email string, cpf string) (*Customer, error) {
	emailVO, err := valueobjects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	cpfVO, err := valueobjects.NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	return &Customer{
		id:    id,
		name:  name,
		email: emailVO,
		cpf:   cpfVO,
	}, nil
}

func (c *Customer) GetId() string {
	return c.id
}

func (c *Customer) GetName() string {
	return c.name
}

func (c *Customer) GetEmail() *valueobjects.Email {
	return c.email
}

func (c *Customer) GetCPF() *valueobjects.CPF {
	return c.cpf
}
