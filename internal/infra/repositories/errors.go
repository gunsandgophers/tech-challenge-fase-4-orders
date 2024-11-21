package repositories

import "errors"

var (
	ErrNotFound = "no rows in result set"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrOrderNotFound    = errors.New("order not found")
)
