package appointment

type AppointmentRepository interface {
	FindById(id uint) (*Appointment, error)
	FindByBarberId(barberId string) ([]*Appointment, error)
	FindByCustomerId(customerId string) ([]*Appointment, error)
	Save(appointment *Appointment) (*Appointment, error)
}
