package valueobjects

type OrderItem struct {
	amount      float64
	quantity    int
	productName string
}

func NewOrderItem(amount float64, quantity int, productName string) *OrderItem {
	return &OrderItem{
		amount: amount,
		quantity: quantity,
		productName: productName,
	}
}

func (i *OrderItem) GetAmount() float64 {
	return i.amount
}

func (i *OrderItem) GetQuantity() int {
	return i.quantity
}

func (i *OrderItem) SetQuatity(quantity int) {
	i.quantity = quantity
}

func (i *OrderItem) GetProductName() string {
	return i.productName
}

func (i *OrderItem) GetTotal() float64 {
	return i.amount * float64(i.quantity)
}
