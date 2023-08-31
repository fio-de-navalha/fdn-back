package application

import (
	"errors"
	"fmt"

	entity "github.com/fio-de-navalha/fdn-back/internal/domain/customer/entities"
	repository "github.com/fio-de-navalha/fdn-back/internal/domain/customer/repositories"
	"github.com/fio-de-navalha/fdn-back/pkg/cryptography"
)

type CustomerServices struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerServices(customerRepository repository.CustomerRepository) *CustomerServices {
	return &CustomerServices{
		customerRepository: customerRepository,
	}
}

func (s *CustomerServices) GetManyCustomers() ([]*entity.Customer, error) {
	users, err := s.customerRepository.FindMany()
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return users, nil
}

func (s *CustomerServices) GetCustomerById(id string) (*entity.Customer, error) {
	user, err := s.customerRepository.FindById(id)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return user, nil
}

func (s *CustomerServices) GetCustomerByPhone(phone string) (*entity.Customer, error) {
	user, err := s.customerRepository.FindByPhone(phone)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return user, nil
}

func (s *CustomerServices) RegisterCustomer(input entity.CustomerInput) error {
	hashedPassword, err := cryptography.HashPassword(input.Password)
	if err != nil {
		return err
	}

	input = entity.CustomerInput{
		Name:     input.Name,
		Phone:    input.Phone,
		Password: hashedPassword,
	}

	customer := entity.NewCustomer(input)
	_, err = s.customerRepository.Create(customer)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return nil
}

func (s *CustomerServices) LoginCustomer(input entity.LoginInput) (*entity.LoginResponse, error) {
	customer, err := s.customerRepository.FindByPhone(input.Phone)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("invalid credentials")
	}

	validPassword := cryptography.ComparePassword(customer.Password, input.Password)
	if !validPassword {
		return nil, errors.New("invalid credentials")
	}

	token, err := cryptography.GenerateToken(customer.ID)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		AccessToken: token,
		Customer:    customer,
	}, nil
}
