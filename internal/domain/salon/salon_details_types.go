package salon

type AddSalonMemberRequest struct {
	ProfessionalId string `json:"professionalId" validate:"required,uuid4"`
	Role           string `json:"role"`
}

type AddSalonAddressRequest struct {
	Address string `json:"address"`
}

type AddSalonContactRequest struct {
	Contact string `json:"contact"`
}
