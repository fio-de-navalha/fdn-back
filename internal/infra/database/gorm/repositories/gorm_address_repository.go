package gorm_repository

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/barber"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormAddressRepository struct {
	db *gorm.DB
}

func NewGormAddressRepository() *gormAddressRepository {
	return &gormAddressRepository{
		db: database.DB,
	}
}

func (r *gormAddressRepository) FindByBarberId(barberId string) ([]*barber.Address, error) {
	var addr []*barber.Address
	res := r.db.Find(&addr, "barber_id = ?", barberId)
	if res.Error != nil {
		return nil, res.Error
	}
	return addr, nil
}

func (r *gormAddressRepository) FindById(id string) (*barber.Address, error) {
	var addr barber.Address
	result := r.db.First(&addr, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &addr, nil
}

func (r *gormAddressRepository) Save(addr *barber.Address) (*barber.Address, error) {
	result := r.db.Save(addr)
	if result.Error != nil {
		return nil, result.Error
	}
	return addr, nil
}

func (r *gormAddressRepository) Delete(addrId string) error {
	var addr []*barber.Address
	res := r.db.Delete(&addr, "id = ?", addrId)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
