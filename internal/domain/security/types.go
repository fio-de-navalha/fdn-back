package security

type SecurityQuestionRequest struct {
	UserId   string `json:"UserId" validate:"required,uuid4"`
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
}
