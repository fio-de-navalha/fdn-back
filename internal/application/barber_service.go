package application

import (
	"errors"
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/barber"
	"github.com/fio-de-navalha/fdn-back/pkg/cryptography"
)

type BarberService struct {
	barberRepository barber.BarberRepository
}

func NewBarberService(barberRepository barber.BarberRepository) *BarberService {
	return &BarberService{
		barberRepository: barberRepository,
	}
}

func (s *BarberService) GetManyBarbers() ([]*barber.Barber, error) {
	log.Println("[application.GetManyBarbers] - Getting many barbers")
	bar, err := s.barberRepository.FindMany()
	if err != nil {
		return nil, err
	}
	return bar, nil
}

func (s *BarberService) GetBarberById(id string) (*barber.BarberResponse, error) {
	log.Println("[application.GetManyBarbers] - Getting barber:", id)
	bar, err := s.barberRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if bar == nil {
		return nil, errors.New("barber not found")
	}
	return &barber.BarberResponse{
		ID:        bar.ID,
		Name:      bar.Name,
		Email:     bar.Email,
		CreatedAt: bar.CreatedAt,
		Services:  bar.Services,
		Products:  bar.Products,
	}, nil
}

func (s *BarberService) GetBarberByEmail(email string) (*barber.Barber, error) {
	log.Println("[application.GetBarberByEmail] - Getting barber by email:", email)
	bar, err := s.barberRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return bar, nil
}

func (s *BarberService) RegisterBarber(input barber.RegisterRequest) error {
	log.Println("[application.RegisterBarber] - Validating barber:", input.Email)
	barberExists, err := s.barberRepository.FindByEmail(input.Email)
	if err != nil {
		return err
	}
	if barberExists != nil {
		return errors.New("barber alredy exists")
	}

	hashedPassword, err := cryptography.HashPassword(input.Password)
	if err != nil {
		return err
	}

	input = barber.RegisterRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	log.Println("[application.RegisterBarber] - Creating barber:", input.Email)
	bar := barber.NewBarber(input)
	_, err = s.barberRepository.Save(bar)
	if err != nil {
		return err
	}
	return nil
}

func (s *BarberService) LoginBarber(input barber.LoginRequest) (*barber.LoginResponse, error) {
	log.Println("[application.LoginBarber] - Validating barber:", input.Email)
	bar, err := s.barberRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if bar == nil {
		return nil, errors.New("invalid credentials")
	}

	validPassword := cryptography.ComparePassword(bar.Password, input.Password)
	if !validPassword {
		return nil, errors.New("invalid credentials")
	}

	log.Println("[application.LoginBarber] - Generating token")
	token, err := cryptography.GenerateToken(bar.ID, "barber")
	if err != nil {
		return nil, err
	}

	barberRes := barber.LoginBarberResponse{
		ID:        bar.ID,
		Name:      bar.Name,
		Email:     bar.Email,
		CreatedAt: bar.CreatedAt,
	}

	return &barber.LoginResponse{
		AccessToken: token,
		Barber:      barberRes,
	}, nil
}
