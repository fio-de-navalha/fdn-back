package database

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"gorm.io/gorm"
)

type gormAddressRepository struct {
	db *gorm.DB
}

func NewGormAddressRepository() *gormAddressRepository {
	return &gormAddressRepository{
		db: DB,
	}
}

func (r *gormAddressRepository) FindBySalonId(salonId string) ([]*salon.Address, error) {
	var addr []*salon.Address
	res := r.db.Find(&addr, "salon_id = ?", salonId)
	if res.Error != nil {
		return nil, res.Error
	}
	return addr, nil
}

func (r *gormAddressRepository) FindById(id string, salonId string) (*salon.Address, error) {
	var addr salon.Address
	result := r.db.First(&addr, "id = ? AND salon_id = ?", id, salonId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &addr, nil
}

func (r *gormAddressRepository) Save(addr *salon.Address) (*salon.Address, error) {
	result := r.db.Save(addr)
	if result.Error != nil {
		return nil, result.Error
	}
	return addr, nil
}

func (r *gormAddressRepository) Delete(addrId string) error {
	var addr []*salon.Address
	res := r.db.Delete(&addr, "id = ?", addrId)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
