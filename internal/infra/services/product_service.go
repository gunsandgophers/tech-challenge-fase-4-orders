package services

import (
	"tech-challenge-fase-1/internal/core/entities"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) FindProductByID(id string) (*entities.Product, error) {

	return entities.RestoreProduct(
		id,
		"Um produto qualquer",
		entities.PRODUCT_CATEGORY_DRINKS,
		1337.0, "Um produto qualquer",
		""), nil
}
