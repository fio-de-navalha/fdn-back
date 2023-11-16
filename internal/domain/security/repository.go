package security

type SecurityQuestionRepository interface {
	FindByUserId(userId string) (*SecurityQuestion, error)
	Save(securityQuestion *SecurityQuestion) (*SecurityQuestion, error)
}
