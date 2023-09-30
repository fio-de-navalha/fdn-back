package application

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
)

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