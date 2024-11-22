package entities

import (
	"testing"

	"github.com/google/uuid"
)

func TestRestoreCustomer(t *testing.T) {
	id := uuid.NewString()
	name := "Customer Silva"
	email := "customer@gunsandgophers.com"
	cpf := "98072798014"
	customer, err := RestoreCustomer(id, name, email, cpf)
	if err != nil {
		t.Errorf("Error create valid customer")
	}
	if customer.GetId() != id {
		t.Errorf("Error Id customer")
	}
	if customer.GetName() != name {
		t.Errorf("Error name customer")
	}
	if customer.GetEmail().Value() != email {
		t.Errorf("Error email customer")
	}
	if customer.GetCPF().Value() != cpf {
		t.Errorf("Error CPF customer")
	}
}
