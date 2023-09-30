package salon

import (
	"github.com/google/uuid"
)

type AddSalonContactRequest struct {
	Contact string `json:"contact"`
}

type Contact struct {
	ID      string `json:"id"`
	SalonId string `json:"salonId"`
	Contact string `json:"contact"`
}

func NewContact(salonId string, contact string) *Contact {
	return &Contact{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Contact: contact,
	}
}
