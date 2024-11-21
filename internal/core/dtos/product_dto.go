package dtos

import "tech-challenge-fase-1/internal/core/entities"

type ProductDTO struct {
	ID          string
	Name        string
	Category    string
	Price       float64
	Description string
	Image       string
}

func NewProductDTOFromEntity(product *entities.Product) *ProductDTO {
	return &ProductDTO{
		ID:          product.GetId(),
		Name:        product.GetName(),
		Category:    product.GetCategory().String(),
		Price:       product.GetPrice(),
		Description: product.GetDescription(),
		Image:       product.GetImage(),
	}
}
