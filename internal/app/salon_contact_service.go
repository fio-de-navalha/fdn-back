package app

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
)

func (s *SalonService) AddSalonContact(salonId, requesterId, contact string) error {
	log.Println("[SalonService.AddSalonContact] - Validating permissions to add contact:", contact)
	sal, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.AddSalonContact] - Creating contact")
	newContact := salon.NewContact(sal.ID, contact)
	if _, err := s.contactRepository.Save(newContact); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonContact(salonId, requesterId, contactId, contact string) (*salon.Contact, error) {
	log.Println("[SalonService.UpdateSalonContact] - Validating permissions to update contact:", contactId)
	_, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.UpdateSalonContact] - Validating contact:", contactId)
	cntt, err := s.validateSalonContact(contactId, salonId)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.UpdateSalonContact] - Updating contact:", contactId)
	cntt.Contact = contact
	if _, err := s.contactRepository.Save(cntt); err != nil {
		return nil, err
	}
	return cntt, nil
}

func (s *SalonService) RemoveSalonContact(salonId, requesterId, contactId string) error {
	log.Println("[SalonService.RemoveSalonContact] - Validating permissions to remove contact:", contactId)
	_, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.RemoveSalonContact] - Validating contact:", contactId)
	cntt, err := s.validateSalonContact(contactId, salonId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.RemoveSalonContact] - Deleting contact:", contactId)
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
		return nil, &utils.AppError{
			Code:    constants.CONTACT_NOT_FOUND_ERROR_CODE,
			Message: constants.CONTACT_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return cntt, nil
}
