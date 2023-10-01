package database

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"gorm.io/gorm"
)

type gormPeriodRepository struct {
	db *gorm.DB
}

func NewGormPeriodRepository() *gormPeriodRepository {
	return &gormPeriodRepository{
		db: DB,
	}
}

func (r *gormPeriodRepository) FindBySalonId(salonId string) ([]*salon.Period, error) {
	var p []*salon.Period
	result := r.db.Find(&p, "salon_id = ?", salonId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return p, nil
}

func (r *gormPeriodRepository) FindById(id string, salonId string) (*salon.Period, error) {
	var p salon.Period
	result := r.db.First(&p, "id = ? AND salon_id = ?", id, salonId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &p, nil
}

func (r *gormPeriodRepository) Save(s *salon.Period) (*salon.Period, error) {
	result := r.db.Save(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

func (r *gormPeriodRepository) Delete(id string) error {
	var p []*salon.Period
	res := r.db.Delete(&p, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
