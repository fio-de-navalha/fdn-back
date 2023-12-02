package app

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
	"github.com/fio-de-navalha/fdn-back/internal/domain/security"
	"github.com/fio-de-navalha/fdn-back/internal/infra/encryption"
	"github.com/fio-de-navalha/fdn-back/pkg/errors"
)

type CustomerService struct {
	customerRepository      customer.CustomerRepository
	securityQuestionService SecurityQuestionService
}

func NewCustomerService(
	customerRepository customer.CustomerRepository,
	securityQuestionService SecurityQuestionService,
) *CustomerService {
	return &CustomerService{
		customerRepository:      customerRepository,
		securityQuestionService: securityQuestionService,
	}
}

func (s *CustomerService) GetCustomerById(id string) (*customer.CustomerResponse, error) {
	log.Println("[CustomerService.GetCustomerById] - Getting customer:", id)
	cus, err := s.customerRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if cus == nil {
		return nil, &errors.AppError{
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
	if cus == nil {
		return nil, &errors.AppError{
			Code:    constants.CUSTOMER_NOT_FOUND_ERROR_CODE,
			Message: constants.CUSTOMER_NOT_FOUND_ERROR_MESSAGE,
		}
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

	log.Println("[CustomerService.RegisterCustomer] - Creating customer")
	cus := customer.NewCustomer(customer.RegisterRequest{
		Name:     input.Name,
		Phone:    input.Phone,
		Password: hashedPassword,
	})
	if _, err = s.customerRepository.Save(cus); err != nil {
		return nil, err
	}

	_, err = s.securityQuestionService.SaveSecurityQuestion(security.SecurityQuestionRequest{
		UserId:   cus.ID,
		Question: input.Question,
		Answer:   input.Answer,
	})
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
		return nil, &errors.AppError{
			Code:    constants.INVALID_CREDENTIAL_ERROR_CODE,
			Message: constants.INVALID_CREDENTIAL_ERROR_MESSAGE,
		}
	}

	validPassword := encryption.ComparePassword(cus.Password, input.Password)
	if !validPassword {
		return nil, &errors.AppError{
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

func (s *CustomerService) ForgotPassword(input customer.ForgotPasswordRequest) (*customer.Customer, error) {
	log.Println("[CustomerService.ForgotPassword] - Getting customer by phone:", input.Phone)
	cus, err := s.customerRepository.FindByPhone(input.Phone)
	if err != nil {
		return nil, err
	}
	if cus == nil {
		return nil, &errors.AppError{
			Code:    constants.CUSTOMER_NOT_FOUND_ERROR_CODE,
			Message: constants.CUSTOMER_NOT_FOUND_ERROR_MESSAGE,
		}
	}

	log.Println("[CustomerService.ForgotPassword] - Validating security question")
	sec, err := s.securityQuestionService.GetByUserId(cus.ID)
	if err != nil {
		return nil, err
	}
	log.Println(sec)

	if sec.Question != input.Question {
		return nil, &errors.AppError{
			Code:    constants.SECURITY_QUESTION_ANSWER_INVALID_ERROR_CODE,
			Message: constants.SECURITY_QUESTION_ANSWER_INVALID_ERROR_MESSAGE,
		}
	}

	if sec.Answer != input.Answer {
		return nil, &errors.AppError{
			Code:    constants.SECURITY_QUESTION_ANSWER_INVALID_ERROR_CODE,
			Message: constants.SECURITY_QUESTION_ANSWER_INVALID_ERROR_MESSAGE,
		}
	}

	return cus, nil
}

func (s *CustomerService) UpdateCustomerPassword(phone string, password string) (*customer.Customer, error) {
	log.Println("[CustomerService.UpdateCustomerPassword] - Getting customer by phone:", phone)
	cus, err := s.customerRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if cus == nil {
		return nil, &errors.AppError{
			Code:    constants.CUSTOMER_NOT_FOUND_ERROR_CODE,
			Message: constants.CUSTOMER_NOT_FOUND_ERROR_MESSAGE,
		}
	}

	hashedPassword, err := encryption.HashPassword(password)
	if err != nil {
		return nil, err
	}

	log.Println("[CustomerService.UpdateCustomerPassword] - Updating customer password:", cus.ID)
	cus.Password = hashedPassword
	if _, err = s.customerRepository.Save(cus); err != nil {
		return nil, err
	}

	return cus, nil
}

func (s *CustomerService) validateCustomerByPhone(phone string) (*customer.Customer, error) {
	cust, err := s.customerRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if cust != nil {
		return nil, &errors.AppError{
			Code:    constants.CUSTOMER_ALREADY_EXISTS_ERROR_CODE,
			Message: constants.CUSTOMER_ALREADY_EXISTS_ERROR_MESSAGE,
		}
	}
	return cust, nil
}
