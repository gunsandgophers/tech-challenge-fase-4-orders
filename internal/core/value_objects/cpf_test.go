package valueobjects

import (
	"regexp"
	"testing"
)

func TestNewCPFValid(t *testing.T) {
	digits := "980.727.980-14"
	reg := regexp.MustCompile(`\D`)
	cpf, err := NewCPF(digits)
	if err != nil {
		t.Errorf("Error to create a valid CPF")
	}
	if cpf.Value() != reg.ReplaceAllString(digits, "") {
		t.Errorf("CPF value different from the defined")
	}
}

func TestNewCPFInvalid(t *testing.T) {
	digits := "212.121.121-21"
	if cpf, err := NewCPF(digits); err == nil || cpf != nil {
		t.Errorf("Create a invalid CPF")
	}
}
