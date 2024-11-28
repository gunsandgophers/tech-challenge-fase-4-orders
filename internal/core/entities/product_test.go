package entities

import (
	"testing"

	uuid "github.com/google/uuid"
)

func TestRestoreProduct(t *testing.T) {
	id := uuid.NewString()
	name := "Product 1"
	category := ProductCategory("Meal")
	price := float64(10.4)
	description := "Some description"
	image := "Some image"
	product := RestoreProduct(id, name, category, price, description, image)
	if product.GetId() == "" {
		t.Errorf("Id can't be empty")
	}
	if product.GetName() != name {
		t.Errorf("Error name customer")
	}
	if product.GetCategory().String() != category.String() {
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
