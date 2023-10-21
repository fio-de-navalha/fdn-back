package repositories

import (
	"github.com/fio-de-navalha/fdn-back/domain/salon"
	"github.com/fio-de-navalha/fdn-back/infra/database"
	"gorm.io/gorm"
)

type gormSalonProductRepository struct {
	db *gorm.DB
}

func NewGormSalonProductRepository() *gormSalonProductRepository {
	return &gormSalonProductRepository{
		db: database.DB,
	}
}

func (r *gormSalonProductRepository) FindManyByIds(id []string) ([]*salon.Product, error) {
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

func (r *gormSalonProductRepository) FindById(id string) (*salon.Product, error) {
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

func (r *gormSalonProductRepository) FindBySalonId(salonId string) ([]*salon.Product, error) {
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

func (r *gormSalonProductRepository) Save(p *salon.Product) (*salon.Product, error) {
	result := r.db.Save(p)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}
