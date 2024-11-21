package valueobjects

import (
	"testing"
)

func TestNewEmailValid(t *testing.T) {
	raw := "email@gunsandgophers.com"
	email, err := NewEmail(raw)
	if err != nil {
		t.Errorf("Error to create a valid CPF")
	}
	if email.Value() != raw {
		t.Errorf("Email value different from the defined")
	}
}

func TestNewEmailInvalid(t *testing.T) {
	raw := "email#gunsandgophers.com"
	if email, err := NewEmail(raw); err == nil || email != nil {
		t.Errorf("Create a invalid CPF")
	}
}
