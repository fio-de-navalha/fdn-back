package product

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
