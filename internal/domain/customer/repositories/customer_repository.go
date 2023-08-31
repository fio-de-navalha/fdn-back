package customer

import customer "github.com/fio-de-navalha/fdn-back/internal/domain/customer/entities"

type CustomerRepository interface {
	FindMany() ([]*customer.Customer, error)
	FindById(id string) (*customer.Customer, error)
	FindByPhone(phone string) (*customer.Customer, error)
	Create(customer *customer.Customer) (*customer.Customer, error)
}
