package database

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"gorm.io/gorm"
)

type gormProfessionalRepository struct {
	db *gorm.DB
}

func NewGormProfessionalRepository() *gormProfessionalRepository {
	return &gormProfessionalRepository{
		db: DB,
	}
}

func (r *gormProfessionalRepository) FindMany() ([]*professional.Professional, error) {
	var prof []*professional.Professional
	res := r.db.Find(&prof)
	if res.Error != nil {
		return nil, res.Error
	}
	return prof, nil
}

func (r *gormProfessionalRepository) FindById(id string) (*professional.Professional, error) {
	var prof professional.Professional
	result := r.db.Model(&professional.Professional{}).
		Preload("Services").
		First(&prof, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &prof, nil
}

func (r *gormProfessionalRepository) FindByEmail(email string) (*professional.Professional, error) {
	var prof professional.Professional
	result := r.db.First(&prof, "email = ?", email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &prof, nil
}

func (r *gormProfessionalRepository) Save(prof *professional.Professional) (*professional.Professional, error) {
	result := r.db.Save(prof)
	if result.Error != nil {
		return nil, result.Error
	}
	return prof, nil
}
