package orders

import (
	"errors"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/queries"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderDisplayListUseCase(t *testing.T) {

	mockOrderDisplayListQuery := &queries.MockOrderDisplayListQueryInterface{}

	useCase := NewOrderDisplayListUseCase(mockOrderDisplayListQuery)

	mockOrderDisplayListQuery.On("Execute").Return([]*dtos.OrderDisplayDTO{}, nil).Once()

	list, _ := useCase.Execute()

	assert.Equal(t, list, []*dtos.OrderDisplayDTO{})
}

func TestOrderDisplayListUseCaseWithErr(t *testing.T) {

	mockOrderDisplayListQuery := &queries.MockOrderDisplayListQueryInterface{}

	useCase := NewOrderDisplayListUseCase(mockOrderDisplayListQuery)

	mockOrderDisplayListQuery.On("Execute").Return(nil, errors.New("error")).Once()

	_, err := useCase.Execute()

	assert.EqualError(t, err, "error")
}
