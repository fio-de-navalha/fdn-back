package service

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID            string    `json:"id"`
	BarberId      string    `json:"barberId"`
	Name          string    `json:"name"`
	Price         int32     `json:"price"`
	DurationInMin int32     `json:"durationInMin"`
	IsAvailable   bool      `json:"isAvailable"`
	CreatedAt     time.Time `json:"createdAt"`
}

func NewService(input CreateServiceInput) *Service {
	return &Service{
		ID:            uuid.NewString(),
		Name:          input.Name,
		Price:         input.Price,
		DurationInMin: input.DurationInMin,
		IsAvailable:   true,
		CreatedAt:     time.Now(),
	}
}
