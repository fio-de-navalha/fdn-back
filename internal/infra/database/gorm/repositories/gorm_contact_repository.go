package gorm_repository

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/barber"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormContactRepository struct {
	db *gorm.DB
}

func NewGormContactRepository() *gormContactRepository {
	return &gormContactRepository{
		db: database.DB,
	}
}

func (r *gormContactRepository) FindByBarberId(barberId string) ([]*barber.Contact, error) {
	var cntt []*barber.Contact
	res := r.db.Find(&cntt, "barber_id = ?", barberId)
	if res.Error != nil {
		return nil, res.Error
	}
	return cntt, nil
}

func (r *gormContactRepository) FindById(id string) (*barber.Contact, error) {
	var cntt barber.Contact
	result := r.db.First(&cntt, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &cntt, nil
}

func (r *gormContactRepository) Save(cntt *barber.Contact) (*barber.Contact, error) {
	result := r.db.Save(cntt)
	if result.Error != nil {
		return nil, result.Error
	}
	return cntt, nil
}

func (r *gormContactRepository) Delete(cnttId string) error {
	var cntt []*barber.Contact
	res := r.db.Delete(&cntt, "id = ?", cnttId)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
