package valueobjects

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	email := &Email{value: value}
	if !email.validate() {
		return nil, errors.New("Invalid Email")
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
