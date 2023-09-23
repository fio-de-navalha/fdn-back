package application

import (
	"errors"
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/barber"
	"github.com/fio-de-navalha/fdn-back/internal/infra/cryptography"
)

type BarberService struct {
	barberRepository  barber.BarberRepository
	addressRepository barber.AddressRepository
	contactRepository barber.ContactRepository
}

func NewBarberService(
	barberRepository barber.BarberRepository,
	addressRepository barber.AddressRepository,
	contactRepository barber.ContactRepository,
) *BarberService {
	return &BarberService{
		barberRepository:  barberRepository,
		addressRepository: addressRepository,
		contactRepository: contactRepository,
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
		Addresses: bar.Addresses,
		Contacts:  bar.Contacts,
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

func (s *BarberService) RegisterBarber(input barber.RegisterRequest) (*barber.AuthResponse, error) {
	log.Println("[application.RegisterBarber] - Validating barber:", input.Email)
	barberExists, err := s.barberRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if barberExists != nil {
		return nil, errors.New("barber alredy exists")
	}

	hashedPassword, err := cryptography.HashPassword(input.Password)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	log.Println("[application.RegisterBarber] - Generating token")
	token, err := cryptography.GenerateToken(bar.ID, "barber")
	if err != nil {
		return nil, err
	}

	return &barber.AuthResponse{
		AccessToken: token,
		Barber: barber.AuthBarberResponse{
			ID:        bar.ID,
			Name:      bar.Name,
			Email:     bar.Email,
			CreatedAt: bar.CreatedAt,
		},
	}, nil
}

func (s *BarberService) LoginBarber(input barber.LoginRequest) (*barber.AuthResponse, error) {
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

	barberRes := barber.AuthBarberResponse{
		ID:        bar.ID,
		Name:      bar.Name,
		Email:     bar.Email,
		CreatedAt: bar.CreatedAt,
	}

	return &barber.AuthResponse{
		AccessToken: token,
		Barber:      barberRes,
	}, nil
}

func (s *BarberService) AddBarberAddress(barberId string, address string) error {
	log.Println("[application.AddBarberAddress] - Validating barber:", barberId)
	bar, err := s.barberRepository.FindById(barberId)
	if err != nil {
		return err
	}
	if bar == nil {
		return errors.New("barber not found")
	}

	addr := barber.NewAddress(bar.ID, address)
	if _, err := s.addressRepository.Save(addr); err != nil {
		return err
	}
	return nil
}

func (s *BarberService) UpdateBarberAddress(addressId string, address string) (*barber.Address, error) {
	log.Println("[application.UpdateBarberAddress] - Validating address:", addressId)
	addr, err := s.addressRepository.FindById(addressId)
	if err != nil {
		return nil, err
	}
	if addr == nil {
		return nil, errors.New("address not found")
	}

	addr.Address = address
	if _, err := s.addressRepository.Save(addr); err != nil {
		return nil, err
	}
	return addr, nil
}

func (s *BarberService) RemoveBarberAddress(addressId string) error {
	log.Println("[application.RemoveBarberAddress] - Validating address:", addressId)
	addr, err := s.addressRepository.FindById(addressId)
	if err != nil {
		return err
	}
	if addr == nil {
		return errors.New("address not found")
	}

	if err := s.addressRepository.Delete(addr.ID); err != nil {
		return err
	}
	return nil
}

func (s *BarberService) AddBarberContact(barberId string, contact string) error {
	log.Println("[application.AddBarberContact] - Validating barber:", barberId)
	bar, err := s.barberRepository.FindById(barberId)
	if err != nil {
		return err
	}
	if bar == nil {
		return errors.New("barber not found")
	}

	cntt := barber.NewContact(bar.ID, contact)
	if _, err := s.contactRepository.Save(cntt); err != nil {
		return err
	}

	return nil
}

func (s *BarberService) UpdateBarberContact(contactId string, contact string) (*barber.Contact, error) {
	log.Println("[application.UpdateBarberContact] - Validating contact:", contactId)
	cntt, err := s.contactRepository.FindById(contactId)
	if err != nil {
		return nil, err
	}
	if cntt == nil {
		return nil, errors.New("contact not found")
	}

	cntt.Contact = contact
	if _, err := s.contactRepository.Save(cntt); err != nil {
		return nil, err
	}
	return cntt, nil
}

func (s *BarberService) RemoveBarberContact(contactId string) error {
	log.Println("[application.RemoveBarberContact] - Validating contact:", contactId)
	cntt, err := s.contactRepository.FindById(contactId)
	if err != nil {
		return err
	}
	if cntt == nil {
		return errors.New("contact not found")
	}

	if err := s.contactRepository.Delete(cntt.ID); err != nil {
		return err
	}
	return nil
}
