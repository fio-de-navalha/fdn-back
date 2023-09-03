package service

type ServiceRepository interface {
	FindById(id uint) (*Service, error)
	FindByBarberId(barberId string) ([]*Service, error)
	Save(service *Service) (*Service, error)
}
