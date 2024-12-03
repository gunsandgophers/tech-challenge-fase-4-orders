package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindProductByID(t *testing.T) {

	client := mocks.NewMockHTTPClientInterface(t)

	service := NewProductService(client)

	productID := uuid.NewString()

	body := dtos.ProductResponseDTO{
		Product: dtos.ProductDTO{
			ID:          uuid.NewString(),
			Name:        "Um nome",
			Category:    "DRINKS",
			Price:       13.37,
			Description: "Uma descrição",
			Image:       "url",
		},
	}

	raw, _ := json.Marshal(body)

	HTTPResponse := &http.Response{Body: io.NopCloser(bytes.NewReader(raw))}

	client.On("Do", mock.Anything).Return(HTTPResponse, nil)

	response, err := service.FindProductByID(productID)

	assert.Nil(t, err)
	assert.Equal(t, response.GetId(), body.Product.ID)
}

func TestFindProductByIDWithError(t *testing.T) {

	client := mocks.NewMockHTTPClientInterface(t)

	service := NewProductService(client)

	productID := uuid.NewString()

	client.On("Do", mock.Anything).Return(nil, errors.New("error"))

	response, err := service.FindProductByID(productID)

	assert.Error(t, err)
	assert.Nil(t, response)
}
