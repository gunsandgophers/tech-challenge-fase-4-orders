package controllers

import "fmt"

// TODO: Alter Regiter to Register
type RegiterCustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

type CheckoutRequest struct {
	CustomerId  *string  `json:"customer_id"`
	ProductsIds []string `json:"products_ids"`
}

type PaymentRequest struct {
	OrderId       string `json:"order_id"`
	PaymentStatus string `json:"payment_status"`
}

type PreparationStatusUpdateRequest struct {
	PreparationStatus string `json:"preparation_status"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func errParamCantBeEmpty(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) cant be empty", name, typ)
}

func (r *RegiterCustomerRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.CPF == "" {
		return errParamIsRequired("cpf", "string")
	}
	return nil
}

func (p *ProductRequest) ValidateProduct() error {
	if p.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if p.Category == "" {
		return errParamIsRequired("category", "string")
	}
	if p.Price == 0 {
		return errParamIsRequired("price", "float64")
	}
	return nil
}

func (r *CheckoutRequest) Validate() error {
	if len(r.ProductsIds) == 0 {
		return errParamCantBeEmpty("products_ids", "string")
	}
	return nil
}

func (r *PaymentRequest) Validate() error {
	if len(r.OrderId) == 0 {
		return errParamIsRequired("order_id", "string")
	}
	if len(r.PaymentStatus) == 0 {
		return errParamIsRequired("payment_status", "string")
	}
	return nil
}

func (r *PreparationStatusUpdateRequest) Validate() error {
	if len(r.PreparationStatus) == 0 {
		return errParamIsRequired("preparation_status", "string")
	}
	return nil
}
