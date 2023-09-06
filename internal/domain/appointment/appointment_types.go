package appointment

import "time"

type CreateAppointmentRequest struct {
	BarberId   string    `json:"barberId" validate:"required,uuid4"`
	CustomerId string    `json:"customerId" validate:"required,uuid4"`
	StartsAt   time.Time `json:"startsAt"`
	ServiceIds []uint    `json:"serviceIds"`
	ProductIds []uint    `json:"productIds"`
}
