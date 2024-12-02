package controllers

import "fmt"

type CheckoutRequest struct {
	CustomerId  *string  `json:"customer_id"`
	ProductsIds []string `json:"products_ids"`
}

type PreparationStatusUpdateRequest struct {
	PreparationStatus string `json:"preparation_status"`
}

func (r *CheckoutRequest) Validate() error {
	if len(r.ProductsIds) == 0 {
		return errParamCantBeEmpty("products_ids", "string")
	}
	return nil
}

func (r *PreparationStatusUpdateRequest) Validate() error {
	if len(r.PreparationStatus) == 0 {
		return errParamIsRequired("preparation_status", "string")
	}
	return nil
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func errParamCantBeEmpty(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) cant be empty", name, typ)
}

