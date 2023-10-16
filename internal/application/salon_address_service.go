package application

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
)

func (s *SalonService) AddSalonAddress(salonId, requesterId, address string) error {
	log.Println("[SalonService.AddSalonAddress] - Validating permissions to add address:", address)
	sal, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.AddSalonAddress] - Creating address")
	newAddr := salon.NewAddress(sal.ID, address)
	if _, err := s.addressRepository.Save(newAddr); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonAddress(salonId, requesterId, addressId, address string) (*salon.Address, error) {
	log.Println("[SalonService.UpdateSalonAddress] - Validating permissions to update address:", addressId)
	_, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.UpdateSalonAddress] - Validating address:", addressId)
	addr, err := s.validateSalonAddress(addressId, salonId)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.UpdateSalonAddress] - Updating address:", addressId)
	addr.Address = address
	if _, err := s.addressRepository.Save(addr); err != nil {
		return nil, err
	}
	return addr, nil
}

func (s *SalonService) RemoveSalonAddress(salonId, requesterId, addressId string) error {
	log.Println("[SalonService.RemoveSalonAddress] - Validating permissions to remove address:", addressId)
	_, err := s.validatePermissionToEditSalon(salonId, requesterId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.RemoveSalonAddress] - Validating address:", addressId)
	addr, err := s.validateSalonAddress(addressId, salonId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.RemoveSalonAddress] - Deleting address:", addressId)
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
		return nil, &utils.AppError{
			Code:    constants.ADDRESS_NOT_FOUND_ERROR_CODE,
			Message: constants.ADDRESS_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return addr, nil
}
