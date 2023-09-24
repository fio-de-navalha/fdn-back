package salon

type SalonRepository interface {
	FindMany() ([]*Salon, error)
	FindById(id string) (*Salon, error)
	Save(salon *Salon) (*Salon, error)
}
