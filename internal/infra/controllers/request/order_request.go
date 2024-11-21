package request

type OpenOrderRequest struct {
	CustomerID *string `json:"customer_id"`
}

func (o *OpenOrderRequest) Validate() error {
	if o.CustomerID != nil && *o.CustomerID == "" {
		return ErrParamIsRequired("customer_id", "string")
	}

	return nil
}

type AddOrderItemRequest struct {
	OrderID   string `json:"order_id"`
	ProductID string    `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (a *AddOrderItemRequest) Validate() error {
	if a.ProductID == "" {
		return ErrParamIsRequired("product_id", "string")
	}

	if a.Quantity == 0 {
		return ErrParamIsRequired("quantity", "int")
	}

	if a.OrderID == "" {
		return ErrParamIsRequired("order_id", "string")
	}

	return nil
}
