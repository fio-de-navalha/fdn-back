package database

import (
	"fmt"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
	"gorm.io/gorm"
)

type gormAppointmentRepository struct {
	db *gorm.DB
}

func NewGormAppointmentRepository() *gormAppointmentRepository {
	return &gormAppointmentRepository{
		db: DB,
	}
}

func (r *gormAppointmentRepository) FindById(id string) (*appointment.Appointment, error) {
	var a appointment.Appointment
	result := r.db.
		Select("id", "barber_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at", "canceled_at").
		Preload("Services").
		Preload("Products").
		Where("id = ?", id).
		First(&a)

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
		Select("id", "barber_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at", "canceled_at").
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
		Select("id", "barber_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at", "canceled_at").
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
	res := r.db.
		Select("id", "barber_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at", "canceled_at").
		Where("starts_at <= ? AND ends_at >= ?", endsAt, startsAt).
		Find(&a)

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

	err := r.db.Create(&services).Error
	if err != nil {
		return nil, err
	}

	if len(products) > 0 {
		err = r.db.Create(&products).Error
		if err != nil {
			return nil, err
		}
	}

	return appo, nil
}

func (r *gormAppointmentRepository) Cancel(appo *appointment.Appointment) (*appointment.Appointment, error) {
	result := r.db.Model(&appo).Update("canceled_at", time.Now())
	fmt.Println(result.Error)
	if result.Error != nil {
		return nil, result.Error
	}
	return appo, nil
}
