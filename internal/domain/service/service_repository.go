package service

type ServiceRepository interface {
	FindManyByIds(ids []uint) ([]*Service, error)
	FindById(id uint) (*Service, error)
	FindByBarberId(barberId string) ([]*Service, error)
	Save(service *Service) (*Service, error)
}
