package application

import (
	"errors"
	"fmt"

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
	bar, err := s.barberRepository.FindMany()
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return bar, nil
}

func (s *BarberService) GetBarberById(id string) (*barber.Barber, error) {
	bar, err := s.barberRepository.FindById(id)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return bar, nil
}

func (s *BarberService) GetBarberByEmail(email string) (*barber.Barber, error) {
	bar, err := s.barberRepository.FindByEmail(email)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return bar, nil
}

func (s *BarberService) RegisterBarber(input barber.BarberInput) error {
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

	input = barber.BarberInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	bar := barber.NewBarber(input)
	_, err = s.barberRepository.Save(bar)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return nil
}

func (s *BarberService) LoginBarber(input barber.LoginInput) (*barber.LoginResponse, error) {
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

	token, err := cryptography.GenerateToken(bar.ID, "barber")
	if err != nil {
		return nil, err
	}

	return &barber.LoginResponse{
		AccessToken: token,
		Barber:      bar,
	}, nil
}
