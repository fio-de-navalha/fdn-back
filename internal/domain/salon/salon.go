package salon

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/google/uuid"
)

type Salon struct {
	ID            string                      `json:"id"`
	Name          string                      `json:"name"`
	OwnerID       string                      `json:"ownerId"`
	CreatedAt     time.Time                   `json:"createdAt"`
	Professionals []professional.Professional `json:"professionals"`
	Addresses     []Address                   `json:"addresses"`
	Contacts      []Contact                   `json:"contacts"`
	Services      []service.Service           `json:"services"`
	Products      []product.Product           `json:"products"`
}

func NewSalon(input CreateSalonRequest) *Salon {
	return &Salon{
		ID:        uuid.NewString(),
		Name:      input.Name,
		OwnerID:   input.OwnerID,
		CreatedAt: time.Now(),
	}
}
