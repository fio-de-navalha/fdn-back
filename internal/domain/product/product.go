package product

import "github.com/google/uuid"

type Product struct {
	ID        string `json:"id" gorm:"primaryKey"`
	BarberId  string `json:"barberId"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Available bool   `json:"available"`
}

func NewProduct(input CreateProductRequest) *Product {
	return &Product{
		ID:        uuid.NewString(),
		BarberId:  input.BarberId,
		Name:      input.Name,
		Price:     input.Price,
		Available: true,
	}
}
