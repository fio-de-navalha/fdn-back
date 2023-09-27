package application

import (
	"errors"
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
)

type SalonService struct {
	salonRepository        salon.SalonRepository
	salonMemberRepository  salon.SalonMemberRepository
	addressRepository      salon.AddressRepository
	contactRepository      salon.ContactRepository
	professionalRepository professional.ProfessionalRepository
}

func NewSalonService(
	salonRepository salon.SalonRepository,
	salonMemberRepository salon.SalonMemberRepository,
	addressRepository salon.AddressRepository,
	contactRepository salon.ContactRepository,
	professionalRepository professional.ProfessionalRepository,
) *SalonService {
	return &SalonService{
		salonRepository:        salonRepository,
		salonMemberRepository:  salonMemberRepository,
		addressRepository:      addressRepository,
		contactRepository:      contactRepository,
		professionalRepository: professionalRepository,
	}
}

func (s *SalonService) GetManySalons() ([]*salon.Salon, error) {
	log.Println("[application.GetManySalons] - Getting many salons")
	res, err := s.salonRepository.FindMany()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SalonService) GetSalonById(id string) (*salon.Salon, error) {
	log.Println("[application.GetSalonById] - Getting salon:", id)
	res, err := s.salonRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("salon not found")
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
	log.Println("[application.CreateSalon] - Validating professional:", professionalId)
	prof, err := s.professionalRepository.FindById(professionalId)
	if err != nil {
		return nil, err
	}
	if prof == nil {
		return nil, errors.New("professional not found")
	}

	log.Println("[application.CreateSalon] - Creating salon:", name)
	newSalon := salon.NewSalon(name)
	sal, err := s.salonRepository.Save(newSalon)
	if err != nil {
		return nil, err
	}

	log.Println("[application.CreateSalon] - Adding salon owner:", professionalId)
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
	log.Println("[application.AddSalonMember] - Validating salon:", salonId)
	sal, err := s.salonRepository.FindById(salonId)
	if err != nil {
		return err
	}
	if sal == nil {
		return errors.New("salon not found")
	}

	isRequesterMember := false
	for _, member := range sal.SalonMembers {
		if member.ProfessionalId == requesterId {
			isRequesterMember = true
			break
		}
	}
	if !isRequesterMember {
		return errors.New("permission denied")
	}

	pro, err := s.professionalRepository.FindById(professionalId)
	if err != nil {
		return err
	}
	if pro == nil {
		return errors.New("professional not found")
	}

	for _, member := range sal.SalonMembers {
		if member.ProfessionalId == professionalId {
			return errors.New("professional already exists in salon")
		}
	}

	log.Println("[application.AddSalonMember] - Adding salon member")
	newSalon := salon.NewSalonMember(sal.ID, professionalId, role)
	if _, err := s.salonMemberRepository.Save(newSalon); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) AddSalonAddress(salonId string, address string) error {
	log.Println("[application.AddSalonAddress] - Validating salon:", salonId)
	sal, err := s.salonRepository.FindById(salonId)
	if err != nil {
		return err
	}
	if sal == nil {
		return errors.New("salon not found")
	}

	newAddr := salon.NewAddress(sal.ID, address)
	if _, err := s.addressRepository.Save(newAddr); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonAddress(salonId string, addressId string, address string) (*salon.Address, error) {
	log.Println("[application.UpdateSalonAddress] - Validating address:", addressId)
	addr, err := s.addressRepository.FindById(addressId, salonId)
	if err != nil {
		return nil, err
	}
	if addr == nil {
		return nil, errors.New("address not found")
	}

	addr.Address = address
	if _, err := s.addressRepository.Save(addr); err != nil {
		return nil, err
	}
	return addr, nil
}

func (s *SalonService) RemoveSalonAddress(salonId string, addressId string) error {
	log.Println("[application.RemoveSalonAddress] - Validating address:", addressId)
	addr, err := s.addressRepository.FindById(addressId, salonId)
	if err != nil {
		return err
	}
	if addr == nil {
		return errors.New("address not found")
	}

	if err := s.addressRepository.Delete(addr.ID); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) AddSalonContact(salonId string, contact string) error {
	log.Println("[application.AddSalonContact] - Validating salon:", salonId)
	sal, err := s.salonRepository.FindById(salonId)
	if err != nil {
		return err
	}
	if sal == nil {
		return errors.New("salon not found")
	}

	newContact := salon.NewContact(sal.ID, contact)
	if _, err := s.contactRepository.Save(newContact); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonContact(salonId string, contactId string, contact string) (*salon.Contact, error) {
	log.Println("[application.UpdateSalonContact] - Validating contact:", contactId)
	cntt, err := s.contactRepository.FindById(contactId, salonId)
	if err != nil {
		return nil, err
	}
	if cntt == nil {
		return nil, errors.New("contact not found")
	}

	cntt.Contact = contact
	if _, err := s.contactRepository.Save(cntt); err != nil {
		return nil, err
	}
	return cntt, nil
}

func (s *SalonService) RemoveSalonContact(salonId string, contactId string) error {
	log.Println("[application.RemoveSalonContact] - Validating contact:", contactId)
	cntt, err := s.contactRepository.FindById(contactId, salonId)
	if err != nil {
		return err
	}
	if cntt == nil {
		return errors.New("contact not found")
	}

	if err := s.contactRepository.Delete(cntt.ID); err != nil {
		return err
	}
	return nil
}
