package dtos

type ProductResponseDTO struct {
	Product struct {
		ID          string  `json:"ID"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		Price       float64 `json:"Price"`
		Description string  `json:"Description"`
		Image       string  `json:"Image"`
	} `json:"product"`
}
