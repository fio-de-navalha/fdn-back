package repositories

import (
	"github.com/fio-de-navalha/fdn-back/domain/salon"
	"github.com/fio-de-navalha/fdn-back/infra/database"
	"gorm.io/gorm"
)

type gormSalonContactRepository struct {
	db *gorm.DB
}

func NewGormSalonContactRepository() *gormSalonContactRepository {
	return &gormSalonContactRepository{
		db: database.DB,
	}
}

func (r *gormSalonContactRepository) FindBySalonId(salonId string) ([]*salon.Contact, error) {
	var cntt []*salon.Contact
	res := r.db.Find(&cntt, "salon_id = ?", salonId)
	if res.Error != nil {
		return nil, res.Error
	}
	return cntt, nil
}

func (r *gormSalonContactRepository) FindById(id string, salonId string) (*salon.Contact, error) {
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

func (r *gormSalonContactRepository) Save(cntt *salon.Contact) (*salon.Contact, error) {
	result := r.db.Save(cntt)
	if result.Error != nil {
		return nil, result.Error
	}
	return cntt, nil
}

func (r *gormSalonContactRepository) Delete(cnttId string) error {
	var cntt []*salon.Contact
	res := r.db.Delete(&cntt, "id = ?", cnttId)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
