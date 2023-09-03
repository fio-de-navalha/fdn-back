package gorm_repository

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository() *gormProductRepository {
	return &gormProductRepository{
		db: database.DB,
	}
}

func (r *gormProductRepository) FindById(id uint) (*product.Product, error) {
	var p product.Product
	result := r.db.First(&p, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &p, nil
}

func (r *gormProductRepository) FindByBarberId(barberId string) ([]*product.Product, error) {
	var p []*product.Product
	result := r.db.Find(&p, "barber_id = ?", barberId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return p, nil
}

func (r *gormProductRepository) Save(p *product.Product) (*product.Product, error) {
	result := r.db.Save(p)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}
