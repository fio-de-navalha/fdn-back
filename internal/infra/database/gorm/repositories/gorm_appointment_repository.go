package gorm_repository

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormAppointmentRepository struct {
	db *gorm.DB
}

func NewGormAppointmentRepository() *gormAppointmentRepository {
	return &gormAppointmentRepository{
		db: database.DB,
	}
}

func (r *gormAppointmentRepository) FindById(id string) (*appointment.Appointment, error) {
	var a appointment.Appointment
	result := r.db.
		Preload("Services").
		Preload("Products").
		First(&a, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &a, nil
}

func (r *gormAppointmentRepository) FindByBarberId(barberId string, startsAt time.Time, endsAt time.Time) ([]*appointment.Appointment, error) {
	var a []*appointment.Appointment
	res := r.db.
		Preload("Services").
		Preload("Products").
		Where("barber_id = ? AND starts_at > ? AND ends_at < ?", barberId, startsAt, endsAt).
		Find(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return a, nil
}

func (r *gormAppointmentRepository) FindByCustomerId(customerId string) ([]*appointment.Appointment, error) {
	var a []*appointment.Appointment
	res := r.db.
		Preload("Services").
		Preload("Products").
		Where("customer_id = ?", customerId).
		Find(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return a, nil
}

func (r *gormAppointmentRepository) FindByDates(startsAt time.Time, endsAt time.Time) ([]*appointment.Appointment, error) {
	var a []*appointment.Appointment
	res := r.db.Where("starts_at <= ? AND ends_at >= ?", endsAt, startsAt).Find(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return a, nil
}

func (r *gormAppointmentRepository) Save(
	appo *appointment.Appointment,
	services []*appointment.AppointmentService,
	products []*appointment.AppointmentProduct,
) (*appointment.Appointment, error) {
	result := r.db.Save(&appo)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, service := range services {
		err := r.db.Save(&service).Error
		return nil, err
	}

	for _, product := range products {
		err := r.db.Save(&product).Error
		return nil, err
	}

	return appo, nil
}
