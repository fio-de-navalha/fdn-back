package appointment

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/google/uuid"
)

type Appointment struct {
	ID             string            `json:"id" gorm:"primaryKey"`
	ProfessionalId string            `json:"professionalId"`
	CustomerId     string            `json:"customerId"`
	DurationInMin  int               `json:"durationInMin"`
	TotalAmount    int               `json:"totalAmount"`
	StartsAt       time.Time         `json:"startsAt"`
	EndsAt         time.Time         `json:"endsAt"`
	CreatedAt      time.Time         `json:"createdAt"`
	Services       []service.Service `json:"services" gorm:"many2many:appointment_service;"`
	Products       []product.Product `json:"products" gorm:"many2many:appointment_product;"`
}

func NewAppointment(
	professionalId string,
	customerId string,
	durationInMin int,
	totalAmount int,
	startsAt time.Time,
	endsAt time.Time,
) *Appointment {
	return &Appointment{
		ID:             uuid.NewString(),
		ProfessionalId: professionalId,
		CustomerId:     customerId,
		DurationInMin:  durationInMin,
		TotalAmount:    totalAmount,
		StartsAt:       startsAt,
		EndsAt:         endsAt,
		CreatedAt:      time.Now(),
	}
}
