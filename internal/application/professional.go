package application

import (
	"errors"
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"github.com/fio-de-navalha/fdn-back/internal/infra/cryptography"
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
	log.Println("[application.GetManyProfessionals] - Getting many professionals")
	res, err := s.professionalRepository.FindMany()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProfessionalService) GetProfessionalById(id string) (*professional.ProfessionalResponse, error) {
	log.Println("[application.GetProfessionalById] - Getting professional:", id)
	res, err := s.professionalRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("professional not found")
	}
	return &professional.ProfessionalResponse{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		Services:  res.Services,
	}, nil
}

func (s *ProfessionalService) GetProfessionalByEmail(email string) (*professional.Professional, error) {
	log.Println("[application.GetProfessionalByEmail] - Getting professional by email:", email)
	res, err := s.professionalRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProfessionalService) RegisterProfessional(input professional.RegisterProfessionalRequest) (*professional.AuthResponse, error) {
	log.Println("[application.RegisterProfessional] - Validating professional:", input.Email)
	profExists, err := s.professionalRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if profExists != nil {
		return nil, errors.New("professional alredy exists")
	}

	hashedPassword, err := cryptography.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input = professional.RegisterProfessionalRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	log.Println("[application.RegisterProfessional] - Creating professional:", input.Email)
	res := professional.NewProfessional(input)
	_, err = s.professionalRepository.Save(res)
	if err != nil {
		return nil, err
	}

	log.Println("[application.RegisterProfessional] - Generating token")
	token, err := cryptography.GenerateToken(res.ID, "professional")
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
	log.Println("[application.LoginProfessional] - Validating professional:", input.Email)
	res, err := s.professionalRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("invalid credentials")
	}

	validPassword := cryptography.ComparePassword(res.Password, input.Password)
	if !validPassword {
		return nil, errors.New("invalid credentials")
	}

	log.Println("[application.LoginProfessional] - Generating token")
	token, err := cryptography.GenerateToken(res.ID, "professional")
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