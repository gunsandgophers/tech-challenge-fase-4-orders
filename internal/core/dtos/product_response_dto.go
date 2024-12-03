package dtos

type ProductDTO struct {
	ID          string  `json:"ID"`
	Name        string  `json:"Name"`
	Category    string  `json:"Category"`
	Price       float64 `json:"Price"`
	Description string  `json:"Description"`
	Image       string  `json:"Image"`
}
type ProductResponseDTO struct {
	Product ProductDTO `json:"product"`
}
