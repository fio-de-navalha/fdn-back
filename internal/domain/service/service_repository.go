package service

type ServiceRepository interface {
	FindById(id string) (*Service, error)
	FindByBarberId(barberId string) ([]*Service, error)
	Save(service *Service) (*Service, error)
}
