package valueobjects

import (
	"testing"
)

func TestNewOrderItem(t *testing.T) {
	var amount float64 = 20.50
	var quantity int = 3
	var productName = "Product 1"
	orderItem := NewOrderItem(amount, 2, productName)
	orderItem.SetQuatity(quantity)
	if orderItem == nil {
		t.Errorf("Error to create a valid order item")
	}
	if orderItem.GetAmount() != amount {
		t.Errorf("Error amount order item")
	}
	if orderItem.GetQuantity() != quantity {
		t.Errorf("Error quantity order item")
	}
	if orderItem.GetProductName() != productName {
		t.Errorf("Error product name order item")
	}
	if orderItem.GetTotal() != amount * float64(quantity) {
		t.Errorf("Error calculate total order item")
	}
}
