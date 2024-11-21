package entities

import (
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
	"testing"

	"github.com/google/uuid"
)

func TestCreateOpenOrderAsAnonymous(t *testing.T) {
	order := CreateOpenOrder(nil)
	if order.GetId() == "" {
		t.Errorf("Order Id can't be empty")
	}
	if order.GetPaymentStatus() != ORDER_PAYMENT_PENDING {
		t.Errorf("Initial Payment Status needs to be PENDING")
	}
	if order.GetPreparationStatus() != ORDER_PREPARATION_AWAITING {
		t.Errorf("Initial Preparation Status needs to be AWAITING")
	}
	if order.GetCustomerId() != nil {
		t.Errorf("CustomerId different from the definition")
	}
	if len(order.GetItems()) != 0 {
		t.Errorf("New order items must to be empty")
	}
}

func TestCreateOpenOrderWithCustomer(t *testing.T) {
	customerId := "customer-id"
	order := CreateOpenOrder(&customerId)
	if order.GetId() == "" {
		t.Errorf("Order Id can't be empty")
	}
	if order.GetPaymentStatus() != ORDER_PAYMENT_PENDING {
		t.Errorf("Initial Payment Status needs to be PENDING")
	}
	if order.GetPreparationStatus() != ORDER_PREPARATION_AWAITING {
		t.Errorf("Initial Preparation Status needs to be AWAITING")
	}
	if order.GetCustomerId() != &customerId {
		t.Errorf("CustomerId different from the definition")
	}
	if len(order.GetItems()) != 0 {
		t.Errorf("New order items must to be empty")
	}
}

func TestRestoreOrderWithItems(t *testing.T) {
	customerId := "customer-id"
	items := []*valueobjects.OrderItem{
		valueobjects.NewOrderItem(34.9, 3, "Product 1"),
		valueobjects.NewOrderItem(10.0, 1, "Product 2"),
		valueobjects.NewOrderItem(20.5, 5, "Product 3"),
	}
	order := RestoreOrder(
		uuid.NewString(),
		&customerId,
		items,
		ORDER_PAYMENT_PENDING,
		ORDER_PREPARATION_AWAITING,
	)
	order.AddItem(CreateProduct(
		"Product 4",
		"Meal",
		10.3,
		"Some Description",
		"Some Image",
	), 2)
	product0 := order.FindOrderItem("Product 0")
	product1 := order.FindOrderItem("Product 1")

	if order.GetId() == "" {
		t.Errorf("Order Id can't be empty")
	}
	if order.GetPaymentStatus() != ORDER_PAYMENT_PENDING {
		t.Errorf("Payment status different from the definition")
	}
	if order.GetPreparationStatus() != ORDER_PREPARATION_AWAITING {
		t.Errorf("Preparation status different from the definition")
	}
	if order.GetCustomerId() != &customerId {
		t.Errorf("CustomerId different from the definition")
	}
	if len(order.GetItems()) != 4 {
		t.Errorf("Order Items len different of 4")
	}
	total := float64((34.9 * 3.0) + (10.0 * 1.0) + (20.5 * 5.0) + (10.3 * 2.0))
	if int32(order.GetTotal() * 100) != int32(total * 100) {
		t.Errorf("Order Items total error")
	}
	if product0 != nil {
		t.Errorf("Product that not exist different from nil")
	}
	if product1.GetProductName() != "Product 1" {
		t.Errorf("Returned a different product")
	}
}
