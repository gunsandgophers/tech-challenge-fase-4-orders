package errors

import "errors"

var (
	ErrInvalidCategory = errors.New("invalid category. The category must be sandwich, sidedishes, drinks or desserts")
	ErrProductNotFound = errors.New("product not found")
	ErrOrderNotAwaitingPayment = errors.New("Order not awaiting payment")
	ErrOrderNotAwaitingPreparation = errors.New("Order not awaiting preparation")
	ErrInvalidPaymentStatus = errors.New("Invalid Payment Status")
	ErrInvalidPreparationStatus = errors.New("Invalid Preparation Status")
	ValidCategories    = map[string]bool{
		"sandwich":   true,
		"sidedishes": true,
		"drinks":     true,
		"desserts":   true,
	}
)
