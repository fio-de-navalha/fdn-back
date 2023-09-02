package application

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
)

type ServiceService struct {
	serviceRepository service.ServiceRepository
}

func NewServiceService(serviceRepository service.ServiceRepository) *ServiceService {
	return &ServiceService{
		serviceRepository: serviceRepository,
	}
}

func (s *ServiceService) GetServicesByBarberId(barberId string) ([]*service.Service, error) {
	res, err := s.serviceRepository.FindByBarberId(barberId)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return res, nil
}

func (s *ServiceService) CreateService(input service.CreateServiceInput) error {
	ser := service.NewService(input)
	_, err := s.serviceRepository.Save(ser)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return nil
}

func (s *ServiceService) UpdateService(serviceId string, input service.UpdateServiceInput) error {
	ser, err := s.serviceRepository.FindById(serviceId)
	if err != nil {
		return err
	}
	if ser == nil {
		return errors.New("service not found")
	}

	updateField := func(dest, source interface{}) {
		if source != nil {
			reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(source).Elem())
		}
	}

	updateField(&ser.Name, input.Name)
	updateField(&ser.Price, input.Price)
	updateField(&ser.DurationInMin, input.DurationInMin)
	updateField(&ser.IsAvailable, input.IsAvailable)

	_, err = s.serviceRepository.Save(ser)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return nil
}
