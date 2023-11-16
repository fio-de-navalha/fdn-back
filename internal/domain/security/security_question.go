package security

import (
	"github.com/google/uuid"
)

type SecurityQuestion struct {
	ID       string `json:"id" gorm:"primaryKey"`
	UserId   string `json:"userId"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func NewSecurityQuestion(input SecurityQuestionRequest) *SecurityQuestion {
	return &SecurityQuestion{
		ID:       uuid.NewString(),
		UserId:   input.UserId,
		Question: input.Question,
		Answer:   input.Answer,
	}
}
