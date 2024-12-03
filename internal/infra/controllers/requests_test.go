package controllers

import (
	"tech-challenge-fase-1/internal/core/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckoutRequestValidate(t *testing.T) {

	checkoutRequest := CheckoutRequest{
		CustomerId:  nil,
		ProductsIds: []string{"1337"},
	}

	err := checkoutRequest.Validate()

	assert.Nil(t, err)
}

func TestCheckoutRequestValidateWithError(t *testing.T) {

	checkoutRequest := CheckoutRequest{
		CustomerId:  nil,
		ProductsIds: []string{},
	}

	err := checkoutRequest.Validate()

	assert.Error(t, err)
}

func TestPreparationStatusUpdateRequestValidate(t *testing.T) {

	preparationStatusUpdateRequest := PreparationStatusUpdateRequest{
		PreparationStatus: entities.ORDER_PREPARATION_CANCELED.String(),
	}

	err := preparationStatusUpdateRequest.Validate()

	assert.Nil(t, err)
}

func TestPreparationStatusUpdateRequestValidateWithError(t *testing.T) {

	preparationStatusUpdateRequest := PreparationStatusUpdateRequest{
		PreparationStatus: "",
	}

	err := preparationStatusUpdateRequest.Validate()

	assert.Error(t, err)
}
