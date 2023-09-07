package service

type ServiceRepository interface {
	FindManyByIds(ids []string) ([]*Service, error)
	FindById(id string) (*Service, error)
	FindByBarberId(barberId string) ([]*Service, error)
	Save(service *Service) (*Service, error)
}
