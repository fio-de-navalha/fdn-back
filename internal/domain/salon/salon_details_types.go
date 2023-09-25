package salon

type AddSalonMemberRequest struct {
	ProfessionalId string `json:"professionalId"`
	Role           string `json:"role"`
}

type AddSalonAddressRequest struct {
	Address string `json:"address"`
}

type AddSalonContactRequest struct {
	Contact string `json:"contact"`
}
