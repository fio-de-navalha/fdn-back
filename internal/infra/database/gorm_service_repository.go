package database

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"gorm.io/gorm"
)

type gormServiceRepository struct {
	db *gorm.DB
}

func NewGormServiceRepository() *gormServiceRepository {
	return &gormServiceRepository{
		db: DB,
	}
}

func (r *gormServiceRepository) FindManyByIds(id []string) ([]*salon.Service, error) {
	var s []*salon.Service
	result := r.db.Find(&s, "id IN ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return s, nil
}

func (r *gormServiceRepository) FindById(id string) (*salon.Service, error) {
	var s salon.Service
	result := r.db.First(&s, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &s, nil
}

func (r *gormServiceRepository) FindBySalonId(salonId string) ([]*salon.Service, error) {
	var s []*salon.Service
	result := r.db.Find(&s, "salon_id = ?", salonId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return s, nil
}

func (r *gormServiceRepository) Save(s *salon.Service) (*salon.Service, error) {
	result := r.db.Save(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}
