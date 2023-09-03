package gorm_repository

import (
	"fmt"

	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormServiceRepository struct {
	db *gorm.DB
}

func NewGormServiceRepository() *gormServiceRepository {
	return &gormServiceRepository{
		db: database.DB,
	}
}

func (r *gormServiceRepository) FindById(id uint) (*service.Service, error) {
	var s service.Service
	result := r.db.First(&s, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &s, nil
}

func (r *gormServiceRepository) FindByBarberId(barberId string) ([]*service.Service, error) {
	var s []*service.Service
	result := r.db.Find(&s, "barber_id = ?", barberId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return s, nil
}

func (r *gormServiceRepository) Save(s *service.Service) (*service.Service, error) {
	fmt.Println(s)
	result := r.db.Save(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}
