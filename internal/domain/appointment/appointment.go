package appointment

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
)

type Appointment struct {
	ID            uint              `json:"id" gorm:"primaryKey"`
	BarberId      string            `json:"barberId"`
	CustomerId    string            `json:"customerId"`
	DurationInMin int               `json:"durationInMin"`
	StartsAt      time.Time         `json:"startsAt"`
	EndsAt        time.Time         `json:"endsAt"`
	CreatedAt     time.Time         `json:"createdAt"`
	Services      []service.Service `gorm:"many2many:appointment_service;"`
	Products      []product.Product `gorm:"many2many:appointment_product;"`
}

func NewAppointment(inputs CreateAppointmentRequest) *Appointment {
	durations := []int{0, 2}
	durationInMin := calculateDuration(durations)

	return &Appointment{
		BarberId:      inputs.BarberId,
		CustomerId:    inputs.CustomerId,
		DurationInMin: durationInMin,
		StartsAt:      inputs.StartsAt,
		EndsAt:        time.Now(),
		CreatedAt:     time.Now(),
	}
}

func calculateDuration(durations []int) int {
	total := 0
	for i := 0; i < len(durations); i++ {
		total += int(durations[i])
	}
	return total
}
