package application

import (
	"errors"
	"log"
	"mime/multipart"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/image"
	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
)

type ServiceService struct {
	serviceRepository   service.ServiceRepository
	salonService        SalonService
	professionalService ProfessionalService
	imageStorageService image.ImageStorageService
}

func NewServiceService(
	serviceRepository service.ServiceRepository,
	salonService SalonService,
	professionalService ProfessionalService,
	imageStorageService image.ImageStorageService,
) *ServiceService {
	return &ServiceService{
		serviceRepository:   serviceRepository,
		salonService:        salonService,
		professionalService: professionalService,
		imageStorageService: imageStorageService,
	}
}

func (s *ServiceService) GetServicesBySalonId(salonId string) ([]*service.Service, error) {
	log.Println("[ServiceService.GetServicesBySalonId] - Validating salon:", salonId)
	if _, err := s.validateSalon(salonId); err != nil {
		return nil, err
	}

	log.Println("[ServiceService.GetServicesBySalonId] - Getting services from salon:", salonId)
	res, err := s.serviceRepository.FindBySalonId(salonId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ServiceService) CreateService(input service.CreateServiceRequest, file *multipart.FileHeader) error {
	log.Println("[ServiceService.CreateService] - Validating salon:", input.SalonId)
	sal, err := s.validateSalon(input.SalonId)
	if err != nil {
		return err
	}

	log.Println("[ServiceService.CreateService] - Validating professional:", input.ProfessionalId)
	prof, err := s.validateProfessional(input.ProfessionalId)
	if err != nil {
		return err
	}

	if err := s.salonService.validateRequesterPermission(prof.ID, sal.SalonMembers); err != nil {
		return err
	}

	if file != nil {
		log.Println("[ServiceService.CreateService] - Uploading file")
		file.Filename = constants.FilePrefix + file.Filename
		res, err := s.imageStorageService.UploadImage(file)
		if err != nil {
			return err
		}

		input.ImageId = res.ID
		input.ImageUrl = res.Urls[0]
	}

	log.Println("[ServiceService.CreateService] - Creating service")
	newService := service.NewService(input)
	if _, err = s.serviceRepository.Save(newService); err != nil {
		return err
	}
	return nil
}

func (s *ServiceService) UpdateService(serviceId string, input service.UpdateServiceRequest, file *multipart.FileHeader) error {
	log.Println("[ServiceService.UpdateService] - Validating salon:", input.SalonId)
	sal, err := s.validateSalon(input.SalonId)
	if err != nil {
		return err
	}

	log.Println("[ServiceService.UpdateService] - Validating professional:", input.ProfessionalId)
	prof, err := s.validateProfessional(input.ProfessionalId)
	if err != nil {
		return err
	}

	if err := s.salonService.validateRequesterPermission(prof.ID, sal.SalonMembers); err != nil {
		return err
	}

	log.Println("[ServiceService.UpdateService] - Validating service:", serviceId)
	ser, err := s.validateService(serviceId)
	if err != nil {
		return err
	}

	if file != nil {
		log.Println("[ServiceService.UpdateService] - Updating image")
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

	log.Println("[ServiceService.UpdateService] - Updating service")
	if _, err = s.serviceRepository.Save(ser); err != nil {
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

func (s *ServiceService) validateSalon(salonId string) (*salon.Salon, error) {
	sal, err := s.salonService.GetSalonById(salonId)
	if err != nil {
		return nil, err
	}
	if sal == nil {
		return nil, errors.New("salon not found")
	}
	return sal, nil
}

func (s *ServiceService) validateProfessional(professionalId string) (*professional.ProfessionalResponse, error) {
	prof, err := s.professionalService.GetProfessionalById(professionalId)
	if err != nil {
		return nil, err
	}
	if prof == nil {
		return nil, errors.New("professional not found")
	}
	return prof, nil
}

func (s *ServiceService) validateService(serviceId string) (*service.Service, error) {
	ser, err := s.serviceRepository.FindById(serviceId)
	if err != nil {
		return nil, err
	}
	if ser == nil {
		return nil, errors.New("service not found")
	}

	return ser, nil
}
