package application

import (
	"errors"
	"fmt"

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
		// TODO: add better error handling
		fmt.Println(err)
	}
	return cus, nil
}

func (s *CustomerService) GetCustomerById(id string) (*customer.Customer, error) {
	cus, err := s.customerRepository.FindById(id)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return cus, nil
}

func (s *CustomerService) GetCustomerByPhone(phone string) (*customer.Customer, error) {
	cus, err := s.customerRepository.FindByPhone(phone)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return cus, nil
}

func (s *CustomerService) RegisterCustomer(input customer.CustomerInput) error {
	hashedPassword, err := cryptography.HashPassword(input.Password)
	if err != nil {
		return err
	}

	input = customer.CustomerInput{
		Name:     input.Name,
		Phone:    input.Phone,
		Password: hashedPassword,
	}

	cus := customer.NewCustomer(input)
	_, err = s.customerRepository.Save(cus)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return nil
}

func (s *CustomerService) LoginCustomer(input customer.LoginInput) (*customer.LoginResponse, error) {
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

	token, err := cryptography.GenerateToken(cus.ID)
	if err != nil {
		return nil, err
	}

	return &customer.LoginResponse{
		AccessToken: token,
		Customer:    cus,
	}, nil
}
