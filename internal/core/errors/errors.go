package errors

import "errors"

var (
	ErrInvalidPreparationStatus    = errors.New("Invalid Preparation Status")
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidCPF   = errors.New("invalid cpf")
)
