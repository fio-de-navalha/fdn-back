package appointment

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/google/uuid"
)

type Appointment struct {
	ID            string              `json:"id" gorm:"primaryKey"`
	BarberId      string            `json:"barberId"`
	CustomerId    string            `json:"customerId"`
	DurationInMin int               `json:"durationInMin"`
	StartsAt      time.Time         `json:"startsAt"`
	EndsAt        time.Time         `json:"endsAt"`
	CreatedAt     time.Time         `json:"createdAt"`
	Services      []service.Service `gorm:"many2many:appointment_service;"`
	Products      []product.Product `gorm:"many2many:appointment_product;"`
}

func NewAppointment(
	barberId string,
	customerId string,
	durationInMin int,
	startsAt time.Time,
	endsAt time.Time,
) *Appointment {
	return &Appointment{
		ID: 		   uuid.NewString(),
		BarberId:      barberId,
		CustomerId:    customerId,
		DurationInMin: durationInMin,
		StartsAt:      startsAt,
		EndsAt:        endsAt,
		CreatedAt:     time.Now(),
	}
}
