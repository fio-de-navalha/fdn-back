package app

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/pkg/errors"
)

type SalonService struct {
	salonRepository       salon.SalonRepository
	salonMemberRepository salon.SalonMemberRepository
	addressRepository     salon.AddressRepository
	contactRepository     salon.ContactRepository
	periodRepository      salon.PeriodRepository
	professionalService   ProfessionalService
}

func NewSalonService(
	salonRepository salon.SalonRepository,
	salonMemberRepository salon.SalonMemberRepository,
	addressRepository salon.AddressRepository,
	contactRepository salon.ContactRepository,
	periodRepository salon.PeriodRepository,
	professionalService ProfessionalService,
) *SalonService {
	return &SalonService{
		salonRepository:       salonRepository,
		salonMemberRepository: salonMemberRepository,
		addressRepository:     addressRepository,
		contactRepository:     contactRepository,
		periodRepository:      periodRepository,
		professionalService:   professionalService,
	}
}

func (s *SalonService) GetManySalons() ([]*salon.Salon, error) {
	log.Println("[SalonService.GetManySalons] - Getting many salons")
	res, err := s.salonRepository.FindMany()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SalonService) GetSalonById(id string) (*salon.Salon, error) {
	slog.Info(fmt.Sprintf("[SalonService.GetSalonById] - Getting salon: %s", id))
	res, err := s.validateSalon(id)
	if err != nil {
		return nil, err
	}
	return &salon.Salon{
		ID:           res.ID,
		Name:         res.Name,
		SalonMembers: res.SalonMembers,
		Addresses:    res.Addresses,
		Contacts:     res.Contacts,
		Periods:      res.Periods,
		Services:     res.Services,
		Products:     res.Products,
	}, nil
}

func (s *SalonService) GetSalonByProfessionalId(professionalId string) (*salon.Salon, error) {
	log.Println("[SalonService.GetSalonByProfessionalId] - Getting salon by professional:", professionalId)
	if _, err := s.professionalService.validateProfessionalById(professionalId); err != nil {
		return nil, err
	}

	salonMember, err := s.salonMemberRepository.FindByProfessionalId(professionalId)
	if err != nil {
		return nil, &errors.AppError{
			Code:    constants.SALON_NOT_FOUND_ERROR_CODE,
			Message: constants.SALON_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	res, err := s.salonRepository.FindById(salonMember.SalonId)
	if err != nil {
		return nil, &errors.AppError{
			Code:    constants.SALON_NOT_FOUND_ERROR_CODE,
			Message: constants.SALON_NOT_FOUND_ERROR_MESSAGE,
		}
	}

	return &salon.Salon{
		ID:           res.ID,
		Name:         res.Name,
		SalonMembers: res.SalonMembers,
		Addresses:    res.Addresses,
		Contacts:     res.Contacts,
		Periods:      res.Periods,
		Services:     res.Services,
		Products:     res.Products,
	}, nil
}

func (s *SalonService) CreateSalon(name string, professionalId string) (*salon.Salon, error) {
	log.Println("[SalonService.CreateSalon] - Validating professional:", professionalId)
	if _, err := s.professionalService.validateProfessionalById(professionalId); err != nil {
		return nil, err
	}

	log.Println("[SalonService.CreateSalon] - Creating salon:", name)
	newSalon := salon.NewSalon(name)
	sal, err := s.salonRepository.Save(newSalon)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.CreateSalon] - Adding salon owner:", professionalId)
	newSalonMember := salon.NewSalonMember(sal.ID, professionalId, "owner")
	salMem, err := s.salonMemberRepository.Save(newSalonMember)
	if err != nil {
		return nil, err
	}
	return &salon.Salon{
		ID:        sal.ID,
		Name:      sal.Name,
		CreatedAt: sal.CreatedAt,
		SalonMembers: []salon.SalonMember{
			*salMem,
		},
	}, nil
}

func (s *SalonService) validateSalon(salonId string) (*salon.Salon, error) {
	res, err := s.salonRepository.FindById(salonId)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, &errors.AppError{
			Code:    constants.SALON_NOT_FOUND_ERROR_CODE,
			Message: constants.SALON_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return res, nil
}

func (s *SalonService) validateProfessional(professionalId string) (*professional.Professional, error) {
	res, err := s.professionalService.validateProfessionalById(professionalId)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, &errors.AppError{
			Code:    constants.PROFESSIONAL_NOT_FOUND_ERROR_CODE,
			Message: constants.PROFESSIONAL_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return res, nil
}

func (s *SalonService) validateRequesterPermission(requesterId string, salonMembers []salon.SalonMember) error {
	isRequesterMember := false
	for _, member := range salonMembers {
		if member.ProfessionalId == requesterId {
			isRequesterMember = true
			break
		}
	}
	if !isRequesterMember {
		return &errors.AppError{
			Code:    constants.PERMISSION_DENIED_ERROR_CODE,
			Message: "permisison denied",
		}
	}
	return nil
}

func (s *SalonService) validatePermissionToEditSalon(salonId, requesterId string) (*salon.Salon, error) {
	log.Println("[SalonService.validatePermissionToEditSalon] - Validating salon:", salonId)
	sal, err := s.validateSalon(salonId)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.validatePermissionToEditSalon] - Validating professional:", requesterId)
	prof, err := s.validateProfessional(requesterId)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.validatePermissionToEditSalon] - Validating requester permission:", requesterId)
	if err := s.validateRequesterPermission(prof.ID, sal.SalonMembers); err != nil {
		return nil, err
	}

	return sal, nil
}
