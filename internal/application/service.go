package application

import (
	"errors"
	"mime/multipart"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/image"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
)

type ServiceService struct {
	serviceRepository   service.ServiceRepository
	barberService       BarberService
	imageStorageService image.ImageStorageService
}

func NewServiceService(
	serviceRepository service.ServiceRepository,
	barberService BarberService,
	imageStorageService image.ImageStorageService,
) *ServiceService {
	return &ServiceService{
		serviceRepository:   serviceRepository,
		barberService:       barberService,
		imageStorageService: imageStorageService,
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
		return nil, err
	}
	return res, nil
}

func (s *ServiceService) CreateService(input service.CreateServiceRequest, file *multipart.FileHeader) error {
	barberExists, err := s.barberService.GetBarberById(input.BarberId)
	if err != nil {
		return err
	}
	if barberExists == nil {
		return errors.New("barber not found")
	}

	if file != nil {
		file.Filename = constants.FilePrefix + file.Filename
		res, err := s.imageStorageService.UploadImage(file)
		if err != nil {
			return err
		}

		input.ImageId = res.ID
		input.ImageUrl = res.Urls[0]
	}

	ser := service.NewService(input)
	_, err = s.serviceRepository.Save(ser)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceService) UpdateService(serviceId string, input service.UpdateServiceRequest, file *multipart.FileHeader) error {
	ser, err := s.serviceRepository.FindById(serviceId)
	if err != nil {
		return err
	}
	if ser == nil {
		return errors.New("service not found")
	}

	if file != nil {
		res, err := s.imageStorageService.UpdateImage(ser.ImageId, file)
		if err != nil {
			return err
		}

		ser.ImageId = res.ID
		ser.ImageUrl = res.Urls[0]
	}

	if input.Name != nil {
		ser.Name = *input.Name
	}
	if input.Description != nil {
		ser.Description = *input.Description
	}
	if input.Price != nil {
		ser.Price = *input.Price
	}
	if input.DurationInMin != nil {
		ser.DurationInMin = *input.DurationInMin
	}
	if input.Available != nil {
		ser.Available = *input.Available
	}

	_, err = s.serviceRepository.Save(ser)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceService) getManyServices(serviceIds []string) ([]*service.Service, error) {
	res, err := s.serviceRepository.FindManyByIds(serviceIds)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ServiceService) ValidateServicesAvailability(services []*service.Service) ([]string, int, int) {
	var durationInMin int
	var totalAmount int
	var servicesIdsToSave []string
	for _, v := range services {
		if v.Available {
			totalAmount += v.Price
			durationInMin += v.DurationInMin
			servicesIdsToSave = append(servicesIdsToSave, v.ID)
		}
	}

	return servicesIdsToSave, durationInMin, totalAmount
}
