package application

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
)

type ServiceService struct {
	serviceRepository service.ServiceRepository
	barberService     BarberService
}

func NewServiceService(serviceRepository service.ServiceRepository, barberService BarberService) *ServiceService {
	return &ServiceService{
		serviceRepository: serviceRepository,
		barberService:     barberService,
	}
}

func (s *ServiceService) GetServicesByBarberId(barberId string) ([]*service.Service, error) {
	barberExists, err := s.barberService.GetBarberById(barberId)
	if err != nil {
		return nil, err
	}
	if barberExists == nil {
		return nil, errors.New("barber not found")
	}

	res, err := s.serviceRepository.FindByBarberId(barberId)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return res, nil
}

func (s *ServiceService) CreateService(input service.CreateServiceInput) error {
	barberExists, err := s.barberService.GetBarberById(input.BarberId)
	if err != nil {
		return err
	}
	if barberExists == nil {
		return errors.New("barber not found")
	}

	ser := service.NewService(input)
	_, err = s.serviceRepository.Save(ser)
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
	updateField(&ser.Available, input.Available)

	_, err = s.serviceRepository.Save(ser)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return nil
}
