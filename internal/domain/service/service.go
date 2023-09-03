package service

type Service struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	BarberId      string `json:"barberId"`
	Name          string `json:"name"`
	Price         int32  `json:"price"`
	DurationInMin int32  `json:"durationInMin"`
	Available     bool   `json:"available"`
}

func NewService(input CreateServiceInput) *Service {
	return &Service{
		BarberId:      input.BarberId,
		Name:          input.Name,
		Price:         input.Price,
		DurationInMin: input.DurationInMin,
		Available:     true,
	}
}
