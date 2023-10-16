package appointment

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/google/uuid"
)

type CreateAppointmentRequest struct {
	SalonId        string    `json:"salonId" validate:"required,uuid4"`
	ProfessionalId string    `json:"professionalId" validate:"required,uuid4"`
	CustomerId     string    `json:"customerId" validate:"required,uuid4"`
	StartsAt       time.Time `json:"startsAt" validate:"required"`
	ServiceIds     []string  `json:"serviceIds" validate:"required,min=1"`
	ProductIds     []string  `json:"productIds"`
}

type SaveAppointment struct {
	Appo        Appointment
	ServicesIds []string
	ProductsIds []string
}

type Appointment struct {
	ID             string          `json:"id" gorm:"primaryKey"`
	ProfessionalId string          `json:"professionalId"`
	CustomerId     string          `json:"customerId"`
	DurationInMin  int             `json:"durationInMin"`
	TotalAmount    int             `json:"totalAmount"`
	StartsAt       time.Time       `json:"startsAt"`
	EndsAt         time.Time       `json:"endsAt"`
	CreatedAt      time.Time       `json:"createdAt"`
	Services       []salon.Service `json:"services" gorm:"many2many:appointment_service;"`
	Products       []salon.Product `json:"products" gorm:"many2many:appointment_product;"`
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

type AppointmentRepository interface {
	FindById(id string) (*Appointment, error)
	FindByIdWithJoins(id string) (*Appointment, error)
	FindByProfessionalId(professionalId string, startsAt time.Time, endsAt time.Time) ([]*Appointment, error)
	FindByCustomerId(customerId string) ([]*Appointment, error)
	FindByDates(startsAt time.Time, endsAt time.Time) ([]*Appointment, error)
	Save(appo *Appointment, services []*AppointmentService, products []*AppointmentProduct) (*Appointment, error)
	Cancel(appo *Appointment) (*Appointment, error)
}
