package application

import (
	"errors"

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

func (s *CustomerService) GetManyCustomers() ([]*customer.Customer, error) {
	cus, err := s.customerRepository.FindMany()
	if err != nil {
		return nil, err
	}
	return cus, nil
}

func (s *CustomerService) GetCustomerById(id string) (*customer.CustomerResponse, error) {
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
	cus, err := s.customerRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	return cus, nil
}

func (s *CustomerService) RegisterCustomer(input customer.RegisterRequest) error {
	barberExists, err := s.customerRepository.FindByPhone(input.Phone)
	if err != nil {
		return err
	}
	if barberExists != nil {
		return errors.New("customer alredy exists")
	}

	hashedPassword, err := cryptography.HashPassword(input.Password)
	if err != nil {
		return err
	}

	input = customer.RegisterRequest{
		Name:     input.Name,
		Phone:    input.Phone,
		Password: hashedPassword,
	}

	cus := customer.NewCustomer(input)
	_, err = s.customerRepository.Save(cus)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerService) LoginCustomer(input customer.LoginRequest) (*customer.LoginResponse, error) {
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

	token, err := cryptography.GenerateToken(cus.ID, "customer")
	if err != nil {
		return nil, err
	}

	return &customer.LoginResponse{
		AccessToken: token,
		Customer: customer.LoginCustomerResponse{
			ID:        cus.ID,
			Name:      cus.Name,
			Phone:     cus.Phone,
			CreatedAt: cus.CreatedAt,
		},
	}, nil
}
