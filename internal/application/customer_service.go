package application

import (
	"errors"
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
	"github.com/fio-de-navalha/fdn-back/pkg/cryptography"
)

type CustomerService struct {
	customerRepository customer.CustomerRepository
}

func NewCustomerService(customerRepository customer.CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (s *CustomerService) GetCustomerById(id string) (*customer.CustomerResponse, error) {
	log.Println("[application.GetCustomerById] - Getting customer:", id)
	cus, err := s.customerRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if cus == nil {
		return nil, errors.New("customer not found")
	}
	return &customer.CustomerResponse{
		ID:        cus.ID,
		Name:      cus.Name,
		Phone:     cus.Phone,
		CreatedAt: cus.CreatedAt,
	}, nil
}

func (s *CustomerService) GetCustomerByPhone(phone string) (*customer.Customer, error) {
	log.Println("[application.GetCustomerByPhone] - Getting customer by phone:", phone)
	cus, err := s.customerRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	return cus, nil
}

func (s *CustomerService) RegisterCustomer(input customer.RegisterRequest) (*customer.AuthResponse, error) {
	log.Println("[application.RegisterCustomer] - Getting customer by phone:", input.Phone)
	barberExists, err := s.customerRepository.FindByPhone(input.Phone)
	if err != nil {
		return nil, err
	}
	if barberExists != nil {
		return nil, errors.New("customer alredy exists")
	}

	hashedPassword, err := cryptography.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input = customer.RegisterRequest{
		Name:     input.Name,
		Phone:    input.Phone,
		Password: hashedPassword,
	}

	log.Println("[application.RegisterCustomer] - Creating customer")
	cus := customer.NewCustomer(input)
	_, err = s.customerRepository.Save(cus)
	if err != nil {
		return nil, err
	}

	log.Println("[application.RegisterCustomer] - Generating token")
	token, err := cryptography.GenerateToken(cus.ID, "customer")
	if err != nil {
		return nil, err
	}

	return &customer.AuthResponse{
		AccessToken: token,
		Customer: customer.AuthCustomerResponse{
			ID:        cus.ID,
			Name:      cus.Name,
			Phone:     cus.Phone,
			CreatedAt: cus.CreatedAt,
		},
	}, nil
}

func (s *CustomerService) LoginCustomer(input customer.LoginRequest) (*customer.AuthResponse, error) {
	log.Println("[application.LoginCustomer] - Getting customer by phone:", input.Phone)
	cus, err := s.customerRepository.FindByPhone(input.Phone)
	if err != nil {
		return nil, err
	}
	if cus == nil {
		return nil, errors.New("invalid credentials")
	}

	validPassword := cryptography.ComparePassword(cus.Password, input.Password)
	if !validPassword {
		return nil, errors.New("invalid credentials")
	}

	log.Println("[application.LoginCustomer] - Generating token")
	token, err := cryptography.GenerateToken(cus.ID, "customer")
	if err != nil {
		return nil, err
	}

	return &customer.AuthResponse{
		AccessToken: token,
		Customer: customer.AuthCustomerResponse{
			ID:        cus.ID,
			Name:      cus.Name,
			Phone:     cus.Phone,
			CreatedAt: cus.CreatedAt,
		},
	}, nil
}
