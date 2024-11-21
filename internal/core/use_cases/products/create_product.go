package products

import (
	"strings"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/repositories"
)

type CreateProduct struct {
	productRepository repositories.ProductRepositoryInterface
}

func NewCreateProductUseCase(productRepository repositories.ProductRepositoryInterface) *CreateProduct {
	return &CreateProduct{
		productRepository: productRepository,
	}
}

func (cp *CreateProduct) Execute(productDTO *dtos.ProductDTO) (*dtos.ProductDTO, error) {
	product := entities.CreateProduct(
		productDTO.Name,
		entities.ProductCategory(strings.ToUpper(productDTO.Category)),
		productDTO.Price,
		productDTO.Description,
		productDTO.Image,
	)
	err := cp.productRepository.Insert(product)
	if err != nil {
		return nil, err
	}
	return dtos.NewProductDTOFromEntity(product), nil
}
