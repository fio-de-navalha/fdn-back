package application

import (
	"errors"
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
)

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
