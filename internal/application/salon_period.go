package application

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
)

func (s *SalonService) AddSalonPeriod(salonId string, contact string) error {
	log.Println("[SalonService.AddSalonPeriod] - Validating salon:", salonId)
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

func (s *SalonService) UpdateSalonPeriod(salonId string, contactId string, contact string) (*salon.Contact, error) {
	log.Println("[SalonService.UpdateSalonPeriod] - Validating period:", contactId)
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

func (s *SalonService) RemoveSalonPeriod(salonId string, contactId string) error {
	log.Println("[SalonService.RemoveSalonPeriod] - Validating period:", contactId)
	cntt, err := s.validateSalonContact(contactId, salonId)
	if err != nil {
		return err
	}

	if err := s.contactRepository.Delete(cntt.ID); err != nil {
		return err
	}
	return nil
}