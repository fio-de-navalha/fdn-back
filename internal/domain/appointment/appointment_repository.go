package appointment

import "time"

type AppointmentRepository interface {
	FindById(id uint) (*Appointment, error)
	FindByBarberId(barberId string) ([]*Appointment, error)
	FindByCustomerId(customerId string) ([]*Appointment, error)
	FindByDates(startsAt time.Time, endsAt time.Time) ([]*Appointment, error)
	Save(appointment *Appointment) (*Appointment, error)
}
