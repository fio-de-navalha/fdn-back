package appointment

import "time"

type AppointmentRepository interface {
	FindById(id string) (*Appointment, error)
	FindByBarberId(barberId string, startsAt time.Time, endsAt time.Time) ([]*Appointment, error)
	FindByCustomerId(customerId string) ([]*Appointment, error)
	FindByDates(startsAt time.Time, endsAt time.Time) ([]*Appointment, error)
	Save(appo *Appointment, services []*AppointmentService, products []*AppointmentProduct) (*Appointment, error)
	Cancel(appo *Appointment) (*Appointment, error)
}
