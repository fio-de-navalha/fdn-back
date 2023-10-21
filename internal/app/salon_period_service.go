package app

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
)

func (s *SalonService) GetSalonPeriodByDay(salonId string, day int) (*salon.Period, error) {
	log.Println("[SalonService.GetSalonPeriodByDay] - Validating salon:", salonId)
	if _, err := s.validateSalon(salonId); err != nil {
		return nil, err
	}

	log.Println("[SalonService.GetSalonPeriodByDay] - Getting period")
	period, err := s.periodRepository.FindBySalonAndDay(salonId, day)
	if err != nil {
		return nil, err
	}
	return period, nil
}

func (s *SalonService) AddSalonPeriod(salonId, requesterId string, input salon.AddPeriodRequest) error {
	log.Println("[SalonService.AddSalonPeriod] - Validating salon:", salonId)
	sal, err := s.validateSalon(salonId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.AddSalonPeriod] - Validating professional permission:", requesterId)
	if err := s.validateRequesterPermission(requesterId, sal.SalonMembers); err != nil {
		return err
	}
	if err := utils.ValidateDatetime(input.Open, "open"); err != nil {
		return err
	}
	if err := utils.ValidateDatetime(input.Close, "close"); err != nil {
		return err
	}

	log.Println("[SalonService.AddSalonPeriod] - Creating period")
	newPeriod := salon.NewPeriod(sal.ID, input.Day, input.Open, input.Close)
	if _, err := s.periodRepository.Save(newPeriod); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) UpdateSalonPeriod(salonId, requesterId, periodId string, input salon.UpdatePeriodRequest) (*salon.Period, error) {
	log.Println("[SalonService.UpdateSalonPeriod] - Validating salon:", salonId)
	sal, err := s.validateSalon(salonId)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.UpdateSalonPeriod] - Validating period:", periodId)
	per, err := s.validateSalonPeriod(periodId, salonId)
	if err != nil {
		return nil, err
	}

	log.Println("[SalonService.UpdateSalonPeriod] - Validating professional permission:", requesterId)
	if err := s.validateRequesterPermission(requesterId, sal.SalonMembers); err != nil {
		return nil, err
	}

	if input.Day != nil {
		per.Day = *input.Day
	}
	if input.Open != nil {
		if err := utils.ValidateDatetime(*input.Open, "open"); err != nil {
			return nil, err
		}
		per.Open = *input.Open
	}
	if input.Close != nil {
		if err := utils.ValidateDatetime(*input.Close, "close"); err != nil {
			return nil, err
		}
		per.Close = *input.Close
	}

	log.Println("[SalonService.UpdateSalonPeriod] - Updating period:", periodId)
	if _, err := s.periodRepository.Save(per); err != nil {
		return nil, err
	}
	return per, nil
}

func (s *SalonService) RemoveSalonPeriod(salonId, requesterId, periodId string) error {
	log.Println("[SalonService.RemoveSalonPeriod] - Validating salon:", salonId)
	sal, err := s.validateSalon(salonId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.RemoveSalonPeriod] - Validating period:", periodId)
	per, err := s.validateSalonPeriod(periodId, salonId)
	if err != nil {
		return err
	}

	log.Println("[SalonService.RemoveSalonPeriod] - Validating professional permission:", requesterId)
	if err := s.validateRequesterPermission(requesterId, sal.SalonMembers); err != nil {
		return err
	}

	log.Println("[SalonService.RemoveSalonPeriod] - Deleting period:", periodId)
	if err := s.periodRepository.Delete(per.ID); err != nil {
		return err
	}
	return nil
}

func (s *SalonService) validateSalonPeriod(periodId, salonId string) (*salon.Period, error) {
	period, err := s.periodRepository.FindById(periodId, salonId)
	if err != nil {
		return nil, err
	}
	if period == nil {
		return nil, &utils.AppError{
			Code:    constants.PERIOD_NOT_FOUND_ERROR_CODE,
			Message: constants.PERIOD_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return period, nil
}
