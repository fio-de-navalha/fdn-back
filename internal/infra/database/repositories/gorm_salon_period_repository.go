package repositories

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormSalonPeriodRepository struct {
	db *gorm.DB
}

func NewGormSalonPeriodRepository() *gormSalonPeriodRepository {
	return &gormSalonPeriodRepository{
		db: database.DB,
	}
}

func (r *gormSalonPeriodRepository) FindBySalonId(salonId string) ([]*salon.Period, error) {
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

func (r *gormSalonPeriodRepository) FindBySalonAndDay(salonId string, day int) (*salon.Period, error) {
	var p *salon.Period
	result := r.db.First(&p, "salon_id = ? AND day = ?", salonId, day)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return p, nil
}

func (r *gormSalonPeriodRepository) FindById(id string, salonId string) (*salon.Period, error) {
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

func (r *gormSalonPeriodRepository) Save(s *salon.Period) (*salon.Period, error) {
	result := r.db.Save(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

func (r *gormSalonPeriodRepository) Delete(id string) error {
	var p []*salon.Period
	res := r.db.Delete(&p, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
