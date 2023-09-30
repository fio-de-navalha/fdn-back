package salon

import (
	"time"

	"github.com/google/uuid"
)

type AddSalonMemberRequest struct {
	ProfessionalId string `json:"professionalId" validate:"required,uuid4"`
	Role           string `json:"role"`
}

type SalonMember struct {
	ID             string    `json:"id"`
	SalonId        string    `json:"salonId"`
	ProfessionalId string    `json:"professionalId"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"createdAt"`
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
