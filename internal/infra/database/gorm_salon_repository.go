package database

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"gorm.io/gorm"
)

type gormSalonRepository struct {
	db *gorm.DB
}

func NewGormSalonRepository() *gormSalonRepository {
	return &gormSalonRepository{
		db: DB,
	}
}

func (r *gormSalonRepository) FindMany() ([]*salon.Salon, error) {
	var sal []*salon.Salon
	res := r.db.Find(&sal)
	if res.Error != nil {
		return nil, res.Error
	}
	return sal, nil
}

func (r *gormSalonRepository) FindById(id string) (*salon.Salon, error) {
	var sal salon.Salon
	result := r.db.Model(&salon.Salon{}).
		Preload("Professionals").
		Preload("Addresses").
		Preload("Contacts").
		Preload("Services").
		Preload("Products").
		First(&sal, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &sal, nil
}

func (r *gormSalonRepository) Save(sal *salon.Salon) (*salon.Salon, error) {
	result := r.db.Save(sal)
	if result.Error != nil {
		return nil, result.Error
	}
	return sal, nil
}
