package service

import "github.com/google/uuid"

type Service struct {
	ID            string `json:"id" gorm:"primaryKey"`
	SalonId       string `json:"salonId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
	DurationInMin int    `json:"durationInMin"`
	Available     bool   `json:"available"`
	ImageId       string `json:"imageId"`
	ImageUrl      string `json:"imageUrl"`
}

func NewService(input CreateServiceRequest) *Service {
	return &Service{
		ID:            uuid.NewString(),
		SalonId:       input.SalonId,
		Name:          input.Name,
		Description:   input.Description,
		Price:         input.Price,
		DurationInMin: input.DurationInMin,
		Available:     true,
		ImageId:       input.ImageId,
		ImageUrl:      input.ImageUrl,
	}
}
