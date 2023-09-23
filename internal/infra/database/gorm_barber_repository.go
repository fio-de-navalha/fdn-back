package database

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/barber"
	"gorm.io/gorm"
)

type gormBarberRepository struct {
	db *gorm.DB
}

func NewGormBarberRepository() *gormBarberRepository {
	return &gormBarberRepository{
		db: DB,
	}
}

func (r *gormBarberRepository) FindMany() ([]*barber.Barber, error) {
	var barbers []*barber.Barber
	res := r.db.Find(&barbers)
	if res.Error != nil {
		return nil, res.Error
	}
	return barbers, nil
}

func (r *gormBarberRepository) FindById(id string) (*barber.Barber, error) {
	var bar barber.Barber
	result := r.db.Model(&barber.Barber{}).
		Preload("Services").
		Preload("Products").
		Preload("Addresses").
		Preload("Contacts").
		First(&bar, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &bar, nil
}

func (r *gormBarberRepository) FindByEmail(email string) (*barber.Barber, error) {
	var barber barber.Barber
	result := r.db.First(&barber, "email = ?", email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &barber, nil
}

func (r *gormBarberRepository) Save(barber *barber.Barber) (*barber.Barber, error) {
	result := r.db.Save(barber)
	if result.Error != nil {
		return nil, result.Error
	}
	return barber, nil
}
