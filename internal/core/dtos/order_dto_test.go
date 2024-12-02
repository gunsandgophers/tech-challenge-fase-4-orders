package dtos

import (
	"tech-challenge-fase-1/internal/core/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPaymentStatusFromEntity(t *testing.T) {
	//Arrange
	customerID := uuid.NewString()
	order := entities.CreateOpenOrder(&customerID)
	//Act
	dto := NewOrderDTOFromEntity(order)
	//Assert
	assert.NotNil(t, dto)
	assert.Equal(t, order.GetId(), dto.Id)
	assert.Equal(t, order.GetCustomerId(), dto.CustomerId)
	assert.Equal(t, order.GetPreparationStatus().String(), dto.PreparationStatus)
	assert.Equal(t, order.GetTotal(), dto.Total)
}
