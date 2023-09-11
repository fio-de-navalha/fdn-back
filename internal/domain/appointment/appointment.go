package appointment

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/google/uuid"
)

type Appointment struct {
	ID            string            `json:"id" gorm:"primaryKey"`
	BarberId      string            `json:"barberId"`
	CustomerId    string            `json:"customerId"`
	DurationInMin int               `json:"durationInMin"`
	TotalAmount   int               `json:"totalAmount"`
	StartsAt      time.Time         `json:"startsAt"`
	EndsAt        time.Time         `json:"endsAt"`
	CreatedAt     time.Time         `json:"createdAt"`
	Services      []service.Service `json:"services" gorm:"many2many:appointment_service;"`
	Products      []product.Product `json:"products" gorm:"many2many:appointment_product;"`
}

func NewAppointment(
	barberId string,
	customerId string,
	durationInMin int,
	totalAmount int,
	startsAt time.Time,
	endsAt time.Time,
) *Appointment {
	return &Appointment{
		ID:            uuid.NewString(),
		BarberId:      barberId,
		CustomerId:    customerId,
		DurationInMin: durationInMin,
		TotalAmount:   totalAmount,
		StartsAt:      startsAt,
		EndsAt:        endsAt,
		CreatedAt:     time.Now(),
	}
}
