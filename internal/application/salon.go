package application

import (
	"errors"
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
)

type SalonService struct {
	salonRepository       salon.SalonRepository
	salonMemberRepository salon.SalonMemberRepository
	addressRepository     salon.AddressRepository
	contactRepository     salon.ContactRepository
	professionalService   ProfessionalService
}

func NewSalonService(
	salonRepository salon.SalonRepository,
	salonMemberRepository salon.SalonMemberRepository,
	addressRepository salon.AddressRepository,
	contactRepository salon.ContactRepository,
	professionalService ProfessionalService,
) *SalonService {
	return &SalonService{
		salonRepository:       salonRepository,
		salonMemberRepository: salonMemberRepository,
		addressRepository:     addressRepository,
		contactRepository:     contactRepository,
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
	log.Println("[SalonService.GetSalonById] - Getting salon:", id)
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
		Services:     res.Services,
		Products:     res.Products,
	}, nil
}

func (s *SalonService) CreateSalon(name string, professionalId string) (*salon.Salon, error) {
	log.Println("[SalonService.CreateSalon] - Validating professional:", professionalId)
	if _, err := s.professionalService.ValidateProfessionalById(professionalId); err != nil {
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

func (s *SalonService) AddSalonMember(salonId string, professionalId string, role string, requesterId string) error {
	log.Println("[SalonService.AddSalonMember] - Validating salon:", salonId)
	sal, err := s.validateSalon(salonId)
	if err != nil {
		return err
	}

	if err := s.validateRequesterPermission(requesterId, sal.SalonMembers); err != nil {
		return err
	}

	if _, err := s.professionalService.ValidateProfessionalById(professionalId); err != nil {
		return err
	}

	for _, member := range sal.SalonMembers {
		if member.ProfessionalId == professionalId {
			return errors.New("professional already exists in salon")
		}
	}

	log.Println("[SalonService.AddSalonMember] - Adding salon member")
	newSalon := salon.NewSalonMember(sal.ID, professionalId, role)
	if _, err := s.salonMemberRepository.Save(newSalon); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) AddSalonAddress(salonId string, address string) error {
	log.Println("[SalonService.AddSalonAddress] - Validating salon:", salonId)
	sal, err := s.validateSalon(salonId)
	if err != nil {
		return err
	}

	newAddr := salon.NewAddress(sal.ID, address)
	if _, err := s.addressRepository.Save(newAddr); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonAddress(salonId string, addressId string, address string) (*salon.Address, error) {
	log.Println("[SalonService.UpdateSalonAddress] - Validating address:", addressId)
	addr, err := s.validateSalonAddress(addressId, salonId)
	if err != nil {
		return nil, err
	}

	addr.Address = address
	if _, err := s.addressRepository.Save(addr); err != nil {
		return nil, err
	}
	return addr, nil
}

func (s *SalonService) RemoveSalonAddress(salonId string, addressId string) error {
	log.Println("[SalonService.RemoveSalonAddress] - Validating address:", addressId)
	addr, err := s.validateSalonAddress(addressId, salonId)
	if err != nil {
		return err
	}

	if err := s.addressRepository.Delete(addr.ID); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) AddSalonContact(salonId string, contact string) error {
	log.Println("[SalonService.AddSalonContact] - Validating salon:", salonId)
	sal, err := s.validateSalon(salonId)
	if err != nil {
		return err
	}

	newContact := salon.NewContact(sal.ID, contact)
	if _, err := s.contactRepository.Save(newContact); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonContact(salonId string, contactId string, contact string) (*salon.Contact, error) {
	log.Println("[SalonService.UpdateSalonContact] - Validating contact:", contactId)
	cntt, err := s.validateSalonContact(contactId, salonId)
	if err != nil {
		return nil, err
	}

	cntt.Contact = contact
	if _, err := s.contactRepository.Save(cntt); err != nil {
		return nil, err
	}
	return cntt, nil
}

func (s *SalonService) RemoveSalonContact(salonId string, contactId string) error {
	log.Println("[SalonService.RemoveSalonContact] - Validating contact:", contactId)
	cntt, err := s.validateSalonContact(contactId, salonId)
	if err != nil {
		return err
	}

	if err := s.contactRepository.Delete(cntt.ID); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) validateSalon(salonId string) (*salon.Salon, error) {
	res, err := s.salonRepository.FindById(salonId)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("salon not found")
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
		return errors.New("permission denied")
	}

	return nil
}

func (s *SalonService) validateSalonAddress(addressId, salonId string) (*salon.Address, error) {
	addr, err := s.addressRepository.FindById(addressId, salonId)
	if err != nil {
		return nil, err
	}
	if addr == nil {
		return nil, errors.New("address not found")
	}
	return addr, nil
}

func (s *SalonService) validateSalonContact(contactId, salonId string) (*salon.Contact, error) {
	cntt, err := s.contactRepository.FindById(contactId, salonId)
	if err != nil {
		return nil, err
	}
	if cntt == nil {
		return nil, errors.New("contact not found")
	}
	return cntt, nil
}
