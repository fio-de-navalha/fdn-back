package gorm_repository

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository() *gormCustomerRepository {
	return &gormCustomerRepository{
		db: database.DB,
	}
}

func (r *gormCustomerRepository) FindMany() ([]*customer.Customer, error) {
	var customers []*customer.Customer
	res := r.db.Find(&customers)
	if res.Error != nil {
		return nil, res.Error
	}
	return customers, nil
}

func (r *gormCustomerRepository) FindById(id string) (*customer.Customer, error) {
	var customer customer.Customer
	result := r.db.First(&customer, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &customer, nil
}

func (r *gormCustomerRepository) FindByPhone(phone string) (*customer.Customer, error) {
	var customer customer.Customer
	result := r.db.First(&customer, "phone = ?", phone)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &customer, nil
}

func (r *gormCustomerRepository) Create(customer *customer.Customer) (*customer.Customer, error) {
	result := r.db.Create(customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}
