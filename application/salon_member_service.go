package application

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/constants"
	"github.com/fio-de-navalha/fdn-back/domain/salon"
	"github.com/fio-de-navalha/fdn-back/utils"
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
			return &utils.AppError{
				Code:    constants.PROFESSIONAL_ALREADY_EXISTS_IN_SALON_ERROR_CODE,
				Message: constants.PROFESSIONAL_ALREADY_EXISTS_IN_SALON_ERROR_MESSAGE,
			}
		}
	}

	log.Println("[SalonService.AddSalonMember] - Adding salon member")
	newSalon := salon.NewSalonMember(sal.ID, professionalId, role)
	if _, err := s.salonMemberRepository.Save(newSalon); err != nil {
		return err
	}
	return nil
}
