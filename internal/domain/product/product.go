package product

import "github.com/google/uuid"

type Product struct {
	ID        string `json:"id" gorm:"primaryKey"`
	SalonId   string `json:"salonId"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Available bool   `json:"available"`
	ImageId   string `json:"imageId"`
	ImageUrl  string `json:"imageUrl"`
}

func NewProduct(input CreateProductRequest) *Product {
	return &Product{
		ID:        uuid.NewString(),
		SalonId:   input.SalonId,
		Name:      input.Name,
		Price:     input.Price,
		Available: true,
		ImageId:   input.ImageId,
		ImageUrl:  input.ImageUrl,
	}
}
