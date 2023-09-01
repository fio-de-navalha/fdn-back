package service

type ServiceRepository interface {
	FindMany() ([]*Service, error)
	FindById(id string) (*Service, error)
	FindByBarberId(barberId string) ([]*Service, error)
	Create(service *Service) (*Service, error)
	Update(id string, service *Service) (*Service, error)
}
