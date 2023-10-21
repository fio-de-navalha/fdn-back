package repositories

import (
	"fmt"

	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormSalonMemberRepository struct {
	db *gorm.DB
}

func NewGormSalonMemberRepository() *gormSalonMemberRepository {
	return &gormSalonMemberRepository{
		db: database.DB,
	}
}

func (r *gormSalonMemberRepository) FindBySalonId(salonId string) ([]*salon.SalonMember, error) {
	var salMem []*salon.SalonMember
	res := r.db.Find(&salMem, "salon_id = ?", salonId)
	if res.Error != nil {
		return nil, res.Error
	}
	return salMem, nil
}

func (r *gormSalonMemberRepository) FindById(id string, salonId string) (*salon.SalonMember, error) {
	var salMem salon.SalonMember
	result := r.db.First(&salMem, "id = ? AND salon_id = ?", id, salonId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &salMem, nil
}

func (r *gormSalonMemberRepository) FindByProfessionalId(professionalIdid string) (*salon.SalonMember, error) {
	var salMem salon.SalonMember
	result := r.db.First(&salMem, "professional_id = ?", professionalIdid)

	if result.Error != nil {
		fmt.Println(result.Error)
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &salMem, nil
}

func (r *gormSalonMemberRepository) Save(salMem *salon.SalonMember) (*salon.SalonMember, error) {
	result := r.db.Save(salMem)
	if result.Error != nil {
		return nil, result.Error
	}
	return salMem, nil
}

func (r *gormSalonMemberRepository) Delete(salMemId string) error {
	var salMem []*salon.SalonMember
	res := r.db.Delete(&salMem, "id = ?", salMemId)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
