package database

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository() *gormProductRepository {
	return &gormProductRepository{
		db: DB,
	}
}

func (r *gormProductRepository) FindManyByIds(id []string) ([]*salon.Product, error) {
	var p []*salon.Product
	result := r.db.Find(&p, "id IN ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return p, nil
}

func (r *gormProductRepository) FindById(id string) (*salon.Product, error) {
	var p salon.Product
	result := r.db.First(&p, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &p, nil
}

func (r *gormProductRepository) FindBySalonId(salonId string) ([]*salon.Product, error) {
	var p []*salon.Product
	result := r.db.Find(&p, "salon_id = ?", salonId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return p, nil
}

func (r *gormProductRepository) Save(p *salon.Product) (*salon.Product, error) {
	result := r.db.Save(p)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}
