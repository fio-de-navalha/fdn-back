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
		log.Println("[application.GetManyBarbers] - Error when getting many barbers")
		return nil, err
	}
	log.Println("[application.GetManyBarbers] - Successfully got many barbers")
	return bar, nil
}

func (s *BarberService) GetBarberById(id string) (*barber.BarberResponse, error) {
	log.Println("[application.GetManyBarbers] - Getting barber:", id)
	bar, err := s.barberRepository.FindById(id)
	if err != nil {
		log.Println("[application.GetManyBarbers] - Error when getting barber:", id)
		return nil, err
	}
	if bar == nil {
		log.Println("[application.GetManyBarbers] - Barber not found")
		return nil, errors.New("barber not found")
	}
	log.Println("[application.GetManyBarbers] - Successfully got barber:", id)
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
		log.Println("[application.GetBarberByEmail] - Error when getting barber by email:", email)
		return nil, err
	}
	log.Println("[application.GetBarberByEmail] - Successfully got barber by email:", email)
	return bar, nil
}

func (s *BarberService) RegisterBarber(input barber.RegisterRequest) error {
	log.Println("[application.RegisterBarber] - Validating barber:", input.Email)
	barberExists, err := s.barberRepository.FindByEmail(input.Email)
	if err != nil {
		log.Println("[application.RegisterBarber] - Error when getting barber:", input.Email)
		return err
	}
	if barberExists != nil {
		log.Println("[application.RegisterBarber] - Barber already exists:", input.Email)
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
		log.Println("[application.RegisterBarber] - Error when creating barber:", input.Email)
		return err
	}

	log.Println("[application.RegisterBarber] - Successfully created barber:", input.Email)
	return nil
}

func (s *BarberService) LoginBarber(input barber.LoginRequest) (*barber.LoginResponse, error) {
	log.Println("[application.LoginBarber] - Validating barber:", input.Email)
	bar, err := s.barberRepository.FindByEmail(input.Email)
	if err != nil {
		log.Println("[application.LoginBarber] - Error when getting barber:", input.Email)
		return nil, err
	}
	if bar == nil {
		log.Println("[application.LoginBarber] - Invalid credentials")
		return nil, errors.New("invalid credentials")
	}

	validPassword := cryptography.ComparePassword(bar.Password, input.Password)
	if !validPassword {
		log.Println("[application.LoginBarber] - Invalid credentials")
		return nil, errors.New("invalid credentials")
	}

	log.Println("[application.LoginBarber] - Generating token")
	token, err := cryptography.GenerateToken(bar.ID, "barber")
	if err != nil {
		log.Println("[application.LoginBarber] - Error when generating token")
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
