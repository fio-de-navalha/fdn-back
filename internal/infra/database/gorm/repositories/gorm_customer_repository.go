package gorm_repository

import (
	customer "github.com/fio-de-navalha/fdn-back/internal/domain/customer/entities"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository() *GormCustomerRepository {
	return &GormCustomerRepository{
		db: database.DB,
	}
}

func (r *GormCustomerRepository) FindMany() ([]*customer.Customer, error) {
	var customers []*customer.Customer
	res := r.db.Find(&customers)
	if res.Error != nil {
		return nil, res.Error
	}
	return customers, nil
}

func (r *GormCustomerRepository) FindById(id string) (*customer.Customer, error) {
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

func (r *GormCustomerRepository) FindByPhone(phone string) (*customer.Customer, error) {
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

func (r *GormCustomerRepository) Create(customer *customer.Customer) (*customer.Customer, error) {
	result := r.db.Create(customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}
