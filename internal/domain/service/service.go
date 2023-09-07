package service

import "github.com/google/uuid"

type Service struct {
	ID            string `json:"id" gorm:"primaryKey"`
	BarberId      string `json:"barberId"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	DurationInMin int    `json:"durationInMin"`
	Available     bool   `json:"available"`
}

func NewService(input CreateServiceInput) *Service {
	return &Service{
		ID:            uuid.NewString(),
		BarberId:      input.BarberId,
		Name:          input.Name,
		Price:         input.Price,
		DurationInMin: input.DurationInMin,
		Available:     true,
	}
}
