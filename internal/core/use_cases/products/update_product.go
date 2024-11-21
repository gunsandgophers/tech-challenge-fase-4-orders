package products

import (
	"strings"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/repositories"
)

type UpdateProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

func NewUpdateProductUseCase(productRepository repositories.ProductRepositoryInterface) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepository: productRepository,
	}
}

func (upuc *UpdateProductUseCase) Execute(productDTO *dtos.ProductDTO) (*dtos.ProductDTO, error) {
	product, err := upuc.productRepository.FindProductByID(productDTO.ID)
	if err != nil {
		return nil, err
	}

	product = updateProductFromDTO(product, productDTO)

	if err := upuc.productRepository.Update(product); err != nil {
		return nil, err
	}

	return dtos.NewProductDTOFromEntity(product), nil
}

func updateProductFromDTO(product *entities.Product, productDTO *dtos.ProductDTO) *entities.Product {
	if productDTO.Name != "" {
		product.SetName(productDTO.Name)
	}
	if productDTO.Category != "" {
		product.SetCategory(entities.ProductCategory(strings.ToUpper(productDTO.Category)))
	}
	if productDTO.Price != 0 {
		product.SetPrice(productDTO.Price)
	}
	if productDTO.Description != "" {
		product.SetDescription(productDTO.Description)
	}
	if productDTO.Image != "" {
		product.SetImage(productDTO.Image)
	}
	return product
}
