package main

import (
	"fmt"
	"log"
	_ "tech-challenge-fase-1/docs"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/infra/database"
	"tech-challenge-fase-1/internal/infra/repositories"
)

type Seeder struct {
	connection         *database.PGXConnectionAdapter
	customerRepository *repositories.CustomerRepositoryDB
	productRepository  *repositories.ProductRepositoryDB
	orderRepository    *repositories.OrderRepositoryDB
}

func newSeeder() *Seeder {
	seeder := &Seeder{}
	seeder.connection = database.NewPGXConnectionAdapter()
	seeder.customerRepository = repositories.NewCustomerRepositoryDB(seeder.connection)
	seeder.productRepository = repositories.NewProductRepositoryDB(seeder.connection)
	seeder.orderRepository = repositories.NewOrderRepositoryDB(seeder.connection)
	return seeder
}

func main() {
	seeder := newSeeder()
	fmt.Println("Running Seeders")
	err := seeder.productRepository.Insert(entities.CreateProduct(
		"Sandwich 1",
		entities.PRODUCT_CATEGORY_SANDWICH,
		30.0,
		"good product",
		"",
	))
	if err  != nil {
		fmt.Println("Erro ao inserir")
		log.Fatal(err)
	}
	seeder.productRepository.Insert(entities.CreateProduct(
		"Sandwich 2",
		entities.PRODUCT_CATEGORY_SANDWICH,
		35.0,
		"good product",
		"",
	))
	fmt.Println("Finish Seeders")

}

