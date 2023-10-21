package app

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
	"github.com/fio-de-navalha/fdn-back/internal/infra/encryption"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
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
	log.Println("[CustomerService.GetCustomerById] - Getting customer:", id)
	cus, err := s.customerRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if cus == nil {
		return nil, &utils.AppError{
			Code:    constants.CUSTOMER_NOT_FOUND_ERROR_CODE,
			Message: constants.CUSTOMER_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return &customer.CustomerResponse{
		ID:        cus.ID,
		Name:      cus.Name,
		Phone:     cus.Phone,
		CreatedAt: cus.CreatedAt,
	}, nil
}

func (s *CustomerService) GetCustomerByPhone(phone string) (*customer.Customer, error) {
	log.Println("[CustomerService.GetCustomerByPhone] - Getting customer by phone:", phone)
	cus, err := s.customerRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	return cus, nil
}

func (s *CustomerService) RegisterCustomer(input customer.RegisterRequest) (*customer.AuthResponse, error) {
	log.Println("[CustomerService.RegisterCustomer] - Getting customer by phone:", input.Phone)
	if _, err := s.validateCustomerByPhone(input.Phone); err != nil {
		return nil, err
	}

	hashedPassword, err := encryption.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input = customer.RegisterRequest{
		Name:     input.Name,
		Phone:    input.Phone,
		Password: hashedPassword,
	}

	log.Println("[CustomerService.RegisterCustomer] - Creating customer")
	cus := customer.NewCustomer(input)
	_, err = s.customerRepository.Save(cus)
	if err != nil {
		return nil, err
	}

	log.Println("[CustomerService.RegisterCustomer] - Generating token")
	token, err := encryption.GenerateToken(cus.ID, "customer")
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
	log.Println("[CustomerService.LoginCustomer] - Getting customer by phone:", input.Phone)
	cus, err := s.customerRepository.FindByPhone(input.Phone)
	if err != nil {
		return nil, err
	}
	if cus == nil {
		return nil, &utils.AppError{
			Code:    constants.INVALID_CREDENTIAL_ERROR_CODE,
			Message: constants.INVALID_CREDENTIAL_ERROR_MESSAGE,
		}
	}

	validPassword := encryption.ComparePassword(cus.Password, input.Password)
	if !validPassword {
		return nil, &utils.AppError{
			Code:    constants.INVALID_CREDENTIAL_ERROR_CODE,
			Message: constants.INVALID_CREDENTIAL_ERROR_MESSAGE,
		}
	}

	log.Println("[CustomerService.LoginCustomer] - Generating token")
	token, err := encryption.GenerateToken(cus.ID, "customer")
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

func (s *CustomerService) validateCustomerByPhone(phone string) (*customer.Customer, error) {
	cust, err := s.customerRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if cust != nil {
		return nil, &utils.AppError{
			Code:    constants.CUSTOMER_ALREADY_EXISTS_ERROR_CODE,
			Message: constants.CUSTOMER_ALREADY_EXISTS_ERROR_MESSAGE,
		}
	}
	return cust, nil
}
