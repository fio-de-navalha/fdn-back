package appointment

import "time"

type CreateAppointmentRequest struct {
	BarberId   string    `json:"barberId" validate:"required,uuid4"`
	CustomerId string    `json:"customerId" validate:"required,uuid4"`
	StartsAt   time.Time `json:"startsAt" validate:"required"`
	ServiceIds []uint    `json:"serviceIds" validate:"required,min=1"`
	ProductIds []uint    `json:"productIds"`
}
