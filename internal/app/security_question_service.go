package app

import (
	"log/slog"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/security"
	"github.com/fio-de-navalha/fdn-back/pkg/errors"
)

type SecurityQuestionService struct {
	securityQuestionRepository security.SecurityQuestionRepository
}

func NewSecurityQuestionService(securityQuestionRepository security.SecurityQuestionRepository) *SecurityQuestionService {
	return &SecurityQuestionService{
		securityQuestionRepository: securityQuestionRepository,
	}
}

func (s *SecurityQuestionService) GetByUserId(userId string) (*security.SecurityQuestion, error) {
	slog.Info("[SecurityQuestionService.GetByUserId] - Getting security question from user: " + userId)
	sec, err := s.securityQuestionRepository.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	if sec == nil {
		return nil, &errors.AppError{
			Code:    constants.SECURITY_QUESTION_NOT_FOUND_ERROR_CODE,
			Message: constants.SECURITY_QUESTION_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return sec, nil
}

func (s *SecurityQuestionService) SaveSecurityQuestion(input security.SecurityQuestionRequest) (*security.SecurityQuestion, error) {
	slog.Info("[SecurityQuestionService.saveSecurityQuestion] - Saving security question for user: " + input.UserId)
	secQues := security.NewSecurityQuestion(input)
	sec, err := s.securityQuestionRepository.Save(secQues)
	if err != nil {
		return nil, err
	}
	return sec, nil
}
