package service

type ServiceRepository interface {
	FindManyByIds(ids []string) ([]*Service, error)
	FindById(id string) (*Service, error)
	FindBySalonId(salonId string) ([]*Service, error)
	Save(service *Service) (*Service, error)
}
