package product

type CreateProductInput struct {
	BarberId  string `json:"barberId" validate:"required,uuid4,min=1"`
	Name      string `json:"name" validate:"required,min=1"`
	Price     int    `json:"price" validate:"required,min=1"`
	Available bool   `json:"available"`
}

type UpdateProductInput struct {
	Name      *string `json:"name"`
	Price     *int    `json:"price"`
	Available *bool   `json:"available"`
}
