package app

import (
	"log/slog"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/pkg/errors"
)

func (s *SalonService) AddSalonContact(salonId, requesterId, contact string) error {
	slog.Info("[SalonService.AddSalonContact] - Validating permissions to add contact: " + contact)
	sal, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return err
	}

	slog.Info("[SalonService.AddSalonContact] - Creating contact")
	newContact := salon.NewContact(sal.ID, contact)
	if _, err := s.contactRepository.Save(newContact); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonContact(salonId, requesterId, contactId, contact string) (*salon.Contact, error) {
	slog.Info("[SalonService.UpdateSalonContact] - Validating permissions to update contact: " + contactId)
	_, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return nil, err
	}

	slog.Info("[SalonService.UpdateSalonContact] - Validating contact: " + contactId)
	cntt, err := s.validateSalonContact(contactId, salonId)
	if err != nil {
		return nil, err
	}

	slog.Info("[SalonService.UpdateSalonContact] - Updating contact: " + contactId)
	cntt.Contact = contact
	if _, err := s.contactRepository.Save(cntt); err != nil {
		return nil, err
	}
	return cntt, nil
}

func (s *SalonService) RemoveSalonContact(salonId, requesterId, contactId string) error {
	slog.Info("[SalonService.RemoveSalonContact] - Validating permissions to remove contact: " + contactId)
	_, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return err
	}

	slog.Info("[SalonService.RemoveSalonContact] - Validating contact: " + contactId)
	cntt, err := s.validateSalonContact(contactId, salonId)
	if err != nil {
		return err
	}

	slog.Info("[SalonService.RemoveSalonContact] - Deleting contact: " + contactId)
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
		return nil, &errors.AppError{
			Code:    constants.CONTACT_NOT_FOUND_ERROR_CODE,
			Message: constants.CONTACT_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return cntt, nil
}
