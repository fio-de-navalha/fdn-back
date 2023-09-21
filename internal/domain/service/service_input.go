package service

type CreateServiceInput struct {
	BarberId      string `json:"barberId" validate:"required,uuid4,min=1"`
	Name          string `json:"name" validate:"required,min=1"`
	Description   string `json:"description"`
	Price         int    `json:"price" validate:"required,min=1"`
	DurationInMin int    `json:"durationInMin" validate:"required,min=1"`
	Available     bool   `json:"available"`
}

type UpdateServiceInput struct {
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Price         *int    `json:"price"`
	DurationInMin *int    `json:"durationInMin"`
	Available     *bool   `json:"available"`
}
