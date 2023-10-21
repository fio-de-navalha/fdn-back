package salon

import (
	"github.com/google/uuid"
)

type CreateProductRequest struct {
	SalonId        string `json:"salonId" validate:"required,uuid4,min=1"`
	ProfessionalId string `json:"professionalId" validate:"required,uuid4,min=1"`
	Name           string `json:"name" validate:"required,min=1"`
	Price          int    `json:"price" validate:"required,min=1"`
	Available      bool   `json:"available"`
	ImageId        string `json:"imageId"`
	ImageUrl       string `json:"imageUrl"`
}

type UpdateProductRequest struct {
	SalonId        string  `json:"salonId" validate:"required,uuid4,min=1"`
	ProfessionalId string  `json:"professionalId" validate:"required,uuid4,min=1"`
	Name           *string `json:"name"`
	Price          *int    `json:"price"`
	Available      *bool   `json:"available"`
	ImageId        *string `json:"imageId"`
	ImageUrl       *string `json:"imageUrl"`
}

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

type ProductRepository interface {
	FindManyByIds(ids []string) ([]*Product, error)
	FindById(id string) (*Product, error)
	FindBySalonId(salonId string) ([]*Product, error)
	Save(product *Product) (*Product, error)
}
