package entities

import (
	"testing"
)

func TestCreateProduct(t *testing.T) {
	name := "Product 1"
	category := "Meal"
	price := float64(10.4)
	description := "Some description"
	image := "Some image"
	product := CreateProduct(name, category, price, description, image)
	if product.GetId() == "" {
		t.Errorf("Id can't be empty")
	}
	if product.GetName() != name {
		t.Errorf("Error name customer")
	}
	if product.GetCategory() != category {
		t.Errorf("Error category customer")
	}
	if product.GetPrice() != price {
		t.Errorf("Error price customer")
	}
	if product.GetDescription() != description {
		t.Errorf("Error description customer")
	}
	if product.GetImage() != image {
		t.Errorf("Error image customer")
	}
}
