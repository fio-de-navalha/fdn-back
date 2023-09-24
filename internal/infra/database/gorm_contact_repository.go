package database

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"gorm.io/gorm"
)

type gormContactRepository struct {
	db *gorm.DB
}

func NewGormContactRepository() *gormContactRepository {
	return &gormContactRepository{
		db: DB,
	}
}

func (r *gormContactRepository) FindBySalonId(salonId string) ([]*salon.Contact, error) {
	var cntt []*salon.Contact
	res := r.db.Find(&cntt, "salon_id = ?", salonId)
	if res.Error != nil {
		return nil, res.Error
	}
	return cntt, nil
}

func (r *gormContactRepository) FindById(id string, salonId string) (*salon.Contact, error) {
	var cntt salon.Contact
	result := r.db.First(&cntt, "id = ? AND salon_id = ?", id, salonId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &cntt, nil
}

func (r *gormContactRepository) Save(cntt *salon.Contact) (*salon.Contact, error) {
	result := r.db.Save(cntt)
	if result.Error != nil {
		return nil, result.Error
	}
	return cntt, nil
}

func (r *gormContactRepository) Delete(cnttId string) error {
	var cntt []*salon.Contact
	res := r.db.Delete(&cntt, "id = ?", cnttId)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
