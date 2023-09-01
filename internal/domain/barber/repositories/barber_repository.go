package barber

import barber "github.com/fio-de-navalha/fdn-back/internal/domain/barber/entities"

type BarberRepository interface {
	FindMany() ([]*barber.Barber, error)
	FindById(id string) (*barber.Barber, error)
	FindByPhone(phone string) (*barber.Barber, error)
	Create(barber *barber.Barber) (*barber.Barber, error)
}
