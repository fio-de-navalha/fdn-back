package repositories

import (
	"github.com/fio-de-navalha/fdn-back/domain/salon"
	"github.com/fio-de-navalha/fdn-back/infra/database"
	"gorm.io/gorm"
)

type gormSalonServiceRepository struct {
	db *gorm.DB
}

func NewGormSalonServiceRepository() *gormSalonServiceRepository {
	return &gormSalonServiceRepository{
		db: database.DB,
	}
}

func (r *gormSalonServiceRepository) FindManyByIds(id []string) ([]*salon.Service, error) {
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

func (r *gormSalonServiceRepository) FindById(id string) (*salon.Service, error) {
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

func (r *gormSalonServiceRepository) FindBySalonId(salonId string) ([]*salon.Service, error) {
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

func (r *gormSalonServiceRepository) Save(s *salon.Service) (*salon.Service, error) {
	result := r.db.Save(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}
