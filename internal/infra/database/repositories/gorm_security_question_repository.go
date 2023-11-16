package repositories

import (
	"github.com/fio-de-navalha/fdn-back/internal/domain/security"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"gorm.io/gorm"
)

type gormSecurityQuestionRepository struct {
	db *gorm.DB
}

func NewGormSecurityQuestionRepository() *gormSecurityQuestionRepository {
	return &gormSecurityQuestionRepository{
		db: database.DB,
	}
}

func (r *gormSecurityQuestionRepository) FindByUserId(userId string) (*security.SecurityQuestion, error) {
	var s *security.SecurityQuestion
	result := r.db.First(&s, "user_id = ?", userId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return s, nil
}

func (r *gormSecurityQuestionRepository) Save(s *security.SecurityQuestion) (*security.SecurityQuestion, error) {
	result := r.db.Save(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}
