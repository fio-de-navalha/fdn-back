package salon

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/google/uuid"
)

type CreateSalonRequest struct {
	Name string `json:"name" validate:"required,min=3,max=30"`
}

type UpdateSalonRequest struct {
	Name *string `json:"name"`
}

type Salon struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	CreatedAt    time.Time         `json:"createdAt"`
	SalonMembers []SalonMember     `json:"salonMembers"`
	Addresses    []Address         `json:"addresses"`
	Contacts     []Contact         `json:"contacts"`
	Periods      []Period          `json:"periods"`
	Services     []service.Service `json:"services"`
	Products     []product.Product `json:"products"`
}

func NewSalon(name string) *Salon {
	return &Salon{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
	}
}
