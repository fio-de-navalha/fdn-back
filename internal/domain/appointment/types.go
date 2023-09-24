package appointment

import "time"

type CreateAppointmentRequest struct {
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
