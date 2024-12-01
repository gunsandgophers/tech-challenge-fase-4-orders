package services

import (
	"encoding/json"
	"io"
	"net/http"
	"tech-challenge-fase-1/internal/core/dtos"
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

	req, err := http.NewRequest("GET", "http://"+config.SERVICE_PRODUCT_URL+"/api/v1/product/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var body dtos.ProductResponseDTO

	err = json.Unmarshal(bodyBytes, &body)

	product := entities.RestoreProduct(
		body.Product.ID,
		body.Product.Name,
		entities.ProductCategory(body.Product.Category),
		body.Product.Price,
		body.Product.Description,
		body.Product.Image,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
