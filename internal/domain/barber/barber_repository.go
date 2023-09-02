package barber

type BarberRepository interface {
	FindMany() ([]*Barber, error)
	FindById(id string) (*Barber, error)
	FindByEmail(email string) (*Barber, error)
	Save(barber *Barber) (*Barber, error)
}
