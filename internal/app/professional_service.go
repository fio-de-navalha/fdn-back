package app

import (
	"log/slog"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"github.com/fio-de-navalha/fdn-back/internal/infra/encryption"
	"github.com/fio-de-navalha/fdn-back/pkg/errors"
)

type ProfessionalService struct {
	professionalRepository professional.ProfessionalRepository
}

func NewProfessionalService(
	professionalRepository professional.ProfessionalRepository,
) *ProfessionalService {
	return &ProfessionalService{
		professionalRepository: professionalRepository,
	}
}

func (s *ProfessionalService) GetManyProfessionals() ([]*professional.Professional, error) {
	slog.Info("[ProfessionalService.GetManyProfessionals] - Getting many professionals")
	res, err := s.professionalRepository.FindMany()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProfessionalService) GetProfessionalById(id string) (*professional.ProfessionalResponse, error) {
	slog.Info("[ProfessionalService.GetProfessionalById] - Getting professional: " + id)
	res, err := s.validateProfessionalById(id)
	if err != nil {
		return nil, err
	}
	return &professional.ProfessionalResponse{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (s *ProfessionalService) GetProfessionalByEmail(email string) (*professional.Professional, error) {
	slog.Info("[ProfessionalService.GetProfessionalByEmail] - Getting professional by email: " + email)
	res, err := s.professionalRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProfessionalService) RegisterProfessional(input professional.RegisterProfessionalRequest) (*professional.AuthResponse, error) {
	slog.Info("[ProfessionalService.RegisterProfessional] - Validating professional: " + input.Email)
	profExists, err := s.professionalRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if profExists != nil {
		return nil, &errors.AppError{
			Code:    constants.PROFESSIONAL_ALREADY_EXISTS_ERROR_CODE,
			Message: constants.PROFESSIONAL_ALREADY_EXISTS_ERROR_MESSAGE,
		}
	}

	hashedPassword, err := encryption.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input = professional.RegisterProfessionalRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	slog.Info("[ProfessionalService.RegisterProfessional] - Creating professional: " + input.Email)
	res := professional.NewProfessional(input)
	_, err = s.professionalRepository.Save(res)
	if err != nil {
		return nil, err
	}

	slog.Info("[ProfessionalService.RegisterProfessional] - Generating token")
	token, err := encryption.GenerateToken(res.ID, "professional")
	if err != nil {
		return nil, err
	}

	return &professional.AuthResponse{
		AccessToken: token,
		Professional: professional.AuthProfessionalResponse{
			ID:        res.ID,
			Name:      res.Name,
			Email:     res.Email,
			CreatedAt: res.CreatedAt,
		},
	}, nil
}

func (s *ProfessionalService) LoginProfessional(input professional.LoginProfessionalRequest) (*professional.AuthResponse, error) {
	slog.Info("[ProfessionalService.LoginProfessional] - Validating professional: " + input.Email)
	res, err := s.professionalRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, &errors.AppError{
			Code:    constants.INVALID_CREDENTIAL_ERROR_CODE,
			Message: constants.INVALID_CREDENTIAL_ERROR_MESSAGE,
		}
	}

	validPassword := encryption.ComparePassword(res.Password, input.Password)
	if !validPassword {
		return nil, &errors.AppError{
			Code:    constants.INVALID_CREDENTIAL_ERROR_CODE,
			Message: constants.INVALID_CREDENTIAL_ERROR_MESSAGE,
		}
	}

	slog.Info("[ProfessionalService.LoginProfessional] - Generating token")
	token, err := encryption.GenerateToken(res.ID, "professional")
	if err != nil {
		return nil, err
	}

	return &professional.AuthResponse{
		AccessToken: token,
		Professional: professional.AuthProfessionalResponse{
			ID:        res.ID,
			Name:      res.Name,
			Email:     res.Email,
			CreatedAt: res.CreatedAt,
		},
	}, nil
}

func (s *ProfessionalService) validateProfessionalById(professionalId string) (*professional.Professional, error) {
	prof, err := s.professionalRepository.FindById(professionalId)
	if err != nil {
		return nil, err
	}
	if prof == nil {
		return nil, &errors.AppError{
			Code:    constants.PROFESSIONAL_NOT_FOUND_ERROR_CODE,
			Message: constants.PROFESSIONAL_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return prof, nil
}
