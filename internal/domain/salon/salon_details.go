package salon

import (
	"time"

	"github.com/google/uuid"
)

type SalonMember struct {
	ID             string    `json:"id"`
	SalonId        string    `json:"salonId"`
	ProfessionalId string    `json:"professionalId"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"createdAt"`
}

type Address struct {
	ID      string `json:"id"`
	SalonId string `json:"salonId"`
	Address string `json:"address"`
}

type Contact struct {
	ID      string `json:"id"`
	SalonId string `json:"salonId"`
	Contact string `json:"contact"`
}

func NewSalonMember(salonId string, professionalId string, role string) *SalonMember {
	return &SalonMember{
		ID:             uuid.NewString(),
		SalonId:        salonId,
		ProfessionalId: professionalId,
		Role:           role,
		CreatedAt:      time.Now(),
	}
}

func NewAddress(salonId string, address string) *Address {
	return &Address{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Address: address,
	}
}

func NewContact(salonId string, contact string) *Contact {
	return &Contact{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Contact: contact,
	}
}
