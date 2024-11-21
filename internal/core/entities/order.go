package entities

import (
	"tech-challenge-fase-1/internal/core/errors"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"

	"github.com/google/uuid"
)

type (
	OrderPaymentStatus     string
	OrderPreparationStatus string
)

func (s OrderPaymentStatus) String() string {
	return string(s)
}

func (s OrderPreparationStatus) String() string {
	return string(s)
}

const (
	ORDER_PAYMENT_PENDING          OrderPaymentStatus = "PENDING"
	ORDER_PAYMENT_PAID             OrderPaymentStatus = "PAID"
	ORDER_PAYMENT_REJECTED         OrderPaymentStatus = "REJECTED"
	ORDER_PAYMENT_AWAITING_PAYMENT OrderPaymentStatus = "AWAITING_PAYMENT"
)

const (
	ORDER_PREPARATION_AWAITING       OrderPreparationStatus = "AWAITING"
	ORDER_PREPARATION_RECEIVED       OrderPreparationStatus = "RECEIVED"
	ORDER_PREPARATION_IN_PREPARARION OrderPreparationStatus = "IN_PREPARATION"
	ORDER_PREPARATION_READY          OrderPreparationStatus = "READY"
	ORDER_PREPARATION_FINISHED       OrderPreparationStatus = "FINISHED"
	ORDER_PREPARATION_CANCELED       OrderPreparationStatus = "CANCELED"
)

func IsValidOrderPreparationStatus(status OrderPreparationStatus) bool {
	switch status {
	case 
		ORDER_PREPARATION_AWAITING,
		ORDER_PREPARATION_RECEIVED,
		ORDER_PREPARATION_IN_PREPARARION,
		ORDER_PREPARATION_READY,
		ORDER_PREPARATION_FINISHED,
		ORDER_PREPARATION_CANCELED:
		return true
	}
	return false
}

type Order struct {
	id                string
	customerId        *string
	items             []*valueobjects.OrderItem
	paymentStatus     OrderPaymentStatus
	preparationStatus OrderPreparationStatus
}

func CreateOpenOrder(customerId *string) *Order {
	return RestoreOrder(
		uuid.NewString(),
		customerId,
		make([]*valueobjects.OrderItem, 0),
		ORDER_PAYMENT_PENDING,
		ORDER_PREPARATION_AWAITING,
	)
}

func RestoreOrder(
	id string,
	customerId *string,
	items []*valueobjects.OrderItem,
	paymentStatus OrderPaymentStatus,
	preparationStatus OrderPreparationStatus,
) *Order {
	return &Order{
		id:                id,
		customerId:        customerId,
		items:             items,
		paymentStatus:     paymentStatus,
		preparationStatus: preparationStatus,
	}
}

func (o *Order) GetId() string {
	return o.id
}

func (o *Order) GetCustomerId() *string {
	return o.customerId
}

func (o *Order) GetItems() []*valueobjects.OrderItem {
	return o.items
}

func (o *Order) GetPaymentStatus() OrderPaymentStatus {
	return o.paymentStatus
}

func (o *Order) AwaitingPayment() {
	o.paymentStatus = ORDER_PAYMENT_AWAITING_PAYMENT
}

func (o *Order) PaymentReceived() {
	o.paymentStatus = ORDER_PAYMENT_PAID
	o.SetPreparationStatus(ORDER_PREPARATION_RECEIVED)
}

func (o *Order) PaymentRejected() {
	o.paymentStatus = ORDER_PAYMENT_REJECTED
	o.SetPreparationStatus(ORDER_PREPARATION_CANCELED)
}

func (o *Order) GetPreparationStatus() OrderPreparationStatus {
	return o.preparationStatus
}

func (o *Order) SetPreparationStatus(status OrderPreparationStatus) error {
	if !IsValidOrderPreparationStatus(status) {
		return errors.ErrInvalidPreparationStatus
	}
	o.preparationStatus = status
	return nil
}

func (o *Order) GetTotal() float64 {
	var total float64
	for _, item := range o.items {
		total = total + item.GetTotal()
	}
	return total
}

func (o *Order) FindOrderItem(productName string) *valueobjects.OrderItem {
	for _, item := range o.items {
		if item.GetProductName() == productName {
			return item
		}
	}
	return nil
}

func (o *Order) AddItem(product *Product, quantity int) {
	amount := product.GetPrice()
	productName := product.GetName()
	item := o.FindOrderItem(productName)
	if item == nil {
		item = valueobjects.NewOrderItem(amount, 0, productName)
		o.items = append(o.items, item)
	}
	quantity = item.GetQuantity() + quantity
	item.SetQuatity(quantity)
}
