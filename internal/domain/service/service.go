package service

import "github.com/google/uuid"

type Service struct {
	ID            string `json:"id" gorm:"primaryKey"`
	BarberId      string `json:"barberId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
	DurationInMin int    `json:"durationInMin"`
	Available     bool   `json:"available"`
}

func NewService(input CreateServiceRequest) *Service {
	return &Service{
		ID:            uuid.NewString(),
		BarberId:      input.BarberId,
		Name:          input.Name,
		Description:   input.Description,
		Price:         input.Price,
		DurationInMin: input.DurationInMin,
		Available:     true,
	}
}
