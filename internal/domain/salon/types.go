package salon

type CreateSalonRequest struct {
	Name string `json:"name" validate:"required,min=3,max=30"`
}

type UpdateSalonRequest struct {
	Name *string `json:"name"`
}
