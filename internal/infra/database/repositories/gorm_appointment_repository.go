package repositories

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
		Select("id", "professional_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at").
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

func (r *gormAppointmentRepository) FindByIdWithJoins(id string) (*appointment.Appointment, error) {
	var a appointment.Appointment
	result := r.db.
		Select("id", "professional_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at").
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

func (r *gormAppointmentRepository) FindByProfessionalId(professionalId string, startsAt time.Time, endsAt time.Time) ([]*appointment.Appointment, error) {
	var a []*appointment.Appointment
	res := r.db.
		Select("id", "professional_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at").
		Preload("Professional").
		Preload("Customer").
		Preload("Services").
		Preload("Products").
		Where("professional_id = ? AND starts_at > ? AND ends_at < ?", professionalId, startsAt, endsAt).
		Find(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return a, nil
}

func (r *gormAppointmentRepository) FindByCustomerId(customerId string, startsAt time.Time) ([]*appointment.Appointment, error) {
	var a []*appointment.Appointment
	res := r.db.
		Select("id", "professional_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at").
		Preload("Professional").
		Preload("Customer").
		Preload("Services").
		Preload("Products").
		Where("customer_id = ? AND starts_at > ?", customerId, startsAt).
		Order("starts_at desc").
		Find(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return a, nil
}

func (r *gormAppointmentRepository) FindByDates(startsAt time.Time, endsAt time.Time) ([]*appointment.Appointment, error) {
	var a []*appointment.Appointment
	res := r.db.
		Select("id", "professional_id", "customer_id", "duration_in_min", "total_amount", "starts_at", "ends_at", "created_at").
		Where("starts_at <= ? AND ends_at > ?", endsAt, startsAt).
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
	result := r.db.Where("appointment_id = ?", appo.ID).Delete(&appointment.AppointmentService{})
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Where("appointment_id = ?", appo.ID).Delete(&appointment.AppointmentProduct{})
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Delete(&appo)
	if result.Error != nil {
		return nil, result.Error
	}
	return appo, nil
}
