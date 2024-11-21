package products

import (
	"tech-challenge-fase-1/internal/core/repositories"
)

func NewDeleteProductUseCase(productRepository repositories.ProductRepositoryInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productRepository: productRepository,
	}
}

type DeleteProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

func (dpc *DeleteProductUseCase) Execute(productID string) error {
	_, err := dpc.productRepository.FindProductByID(productID)
	if err != nil {
		return err
	}

	return dpc.productRepository.Delete(productID)
}
