package salon

type CreateSalonRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=30"`
	OwnerID string `json:"ownerId" validate:"required,uuid4,min=1"`
}

type UpdateSalonRequest struct {
	Name *string `json:"name"`
}

type ManageBarberRequest struct {
	BarberID string `json:"barberId" validate:"required,uuid4,min=1"`
}
