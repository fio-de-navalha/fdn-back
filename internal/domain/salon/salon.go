package salon

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/google/uuid"
)

type Salon struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	CreatedAt    time.Time         `json:"createdAt"`
	SalonMembers []SalonMember     `json:"salonMembers"`
	Addresses    []Address         `json:"addresses"`
	Contacts     []Contact         `json:"contacts"`
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
