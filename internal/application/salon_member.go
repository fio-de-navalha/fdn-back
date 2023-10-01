package application

import (
	"errors"
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
)

func (s *SalonService) AddSalonMember(salonId, professionalId, role, requesterId string) error {
	log.Println("[SalonService.AddSalonMember] - Validating salon:", salonId)
	sal, err := s.validateSalon(salonId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.AddSalonMember] - Validating requester permission:", requesterId)
	if err := s.validateRequesterPermission(requesterId, sal.SalonMembers); err != nil {
		return err
	}

	log.Println("[SalonService.AddSalonMember] - Validating professional:", professionalId)
	if _, err := s.professionalService.validateProfessionalById(professionalId); err != nil {
		return err
	}

	log.Println("[SalonService.AddSalonMember] - Validating if professional is already a member")
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
