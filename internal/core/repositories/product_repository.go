package repositories

import "tech-challenge-fase-1/internal/core/entities"

type ProductRepositoryInterface interface {
	Insert(product *entities.Product) error
	Update(product *entities.Product) error
	FindProductByID(ID string) (*entities.Product, error)
	Delete(ID string) error
	FindProductByCategory(category entities.ProductCategory, page, size int) ([]*entities.Product, error)
	ListProducts() ([]*entities.Product, error)
}
