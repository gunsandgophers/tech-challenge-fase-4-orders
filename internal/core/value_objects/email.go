package valueobjects

import (
	"regexp"
	"tech-challenge-fase-1/internal/core/errors"
)


type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	email := &Email{value: value}
	if !email.validate() {
		return nil, errors.ErrInvalidEmail
	}
	return email, nil
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) validate() bool {
	ok, _ := regexp.MatchString("^(.+)@(.+)$", e.value)
	return ok
}
