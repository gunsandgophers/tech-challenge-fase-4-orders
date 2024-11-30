package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/infra/config"
)

type ProductService struct {
	client *http.Client
}

func NewProductService(client *http.Client) *ProductService {
	return &ProductService{
		client: client,
	}
}

func (p *ProductService) FindProductByID(id string) (*entities.Product, error) {

	resp, err := p.client.Get(fmt.Sprint(config.SERVICE_PRODUCT_URL, "/api/v1/product/", id))

	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var body map[string]any

	err = json.Unmarshal(bodyBytes, &body)

	product := body["product"].(entities.Product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
