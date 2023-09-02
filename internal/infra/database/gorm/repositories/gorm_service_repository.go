package gorm_repository

import (
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

func (r *gormServiceRepository) FindById(id string) (*service.Service, error) {
	var service service.Service
	result := r.db.First(&service, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &service, nil
}

func (r *gormServiceRepository) FindByBarberId(barberId string) ([]*service.Service, error) {
	var services []*service.Service
	result := r.db.Find(&services, "barber_id = ?", barberId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return services, nil
}

func (r *gormServiceRepository) Save(service *service.Service) (*service.Service, error) {
	result := r.db.Save(service)
	if result.Error != nil {
		return nil, result.Error
	}
	return service, nil
}
