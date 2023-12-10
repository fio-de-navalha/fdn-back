package app

import (
	"log/slog"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/pkg/errors"
)

func (s *SalonService) AddSalonAddress(salonId, requesterId, address string) error {
	slog.Info("[SalonService.AddSalonAddress] - Validating permissions to add address: " + address)
	sal, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return err
	}

	slog.Info("[SalonService.AddSalonAddress] - Creating address")
	newAddr := salon.NewAddress(sal.ID, address)
	if _, err := s.addressRepository.Save(newAddr); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonAddress(salonId, requesterId, addressId, address string) (*salon.Address, error) {
	slog.Info("[SalonService.UpdateSalonAddress] - Validating permissions to update address: " + addressId)
	_, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return nil, err
	}

	slog.Info("[SalonService.UpdateSalonAddress] - Validating address: " + addressId)
	addr, err := s.validateSalonAddress(addressId, salonId)
	if err != nil {
		return nil, err
	}

	slog.Info("[SalonService.UpdateSalonAddress] - Updating address: " + addressId)
	addr.Address = address
	if _, err := s.addressRepository.Save(addr); err != nil {
		return nil, err
	}
	return addr, nil
}

func (s *SalonService) RemoveSalonAddress(salonId, requesterId, addressId string) error {
	slog.Info("[SalonService.RemoveSalonAddress] - Validating permissions to remove address: " + addressId)
	_, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return err
	}

	slog.Info("[SalonService.RemoveSalonAddress] - Validating address: " + addressId)
	addr, err := s.validateSalonAddress(addressId, salonId)
	if err != nil {
		return err
	}

	slog.Info("[SalonService.RemoveSalonAddress] - Deleting address: " + addressId)
	if err := s.addressRepository.Delete(addr.ID); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) validateSalonAddress(addressId, salonId string) (*salon.Address, error) {
	addr, err := s.addressRepository.FindById(addressId, salonId)
	if err != nil {
		return nil, err
	}
	if addr == nil {
		return nil, &errors.AppError{
			Code:    constants.ADDRESS_NOT_FOUND_ERROR_CODE,
			Message: constants.ADDRESS_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return addr, nil
}
