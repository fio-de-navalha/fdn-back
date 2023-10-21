package salon

import "github.com/google/uuid"

type CreateServiceRequest struct {
	SalonId        string `json:"salonId" validate:"required,uuid4,min=1"`
	ProfessionalId string `json:"professionalId" validate:"required,uuid4,min=1"`
	Name           string `json:"name" validate:"required,min=1"`
	Description    string `json:"description"`
	Price          int    `json:"price" validate:"required,min=1"`
	DurationInMin  int    `json:"durationInMin" validate:"required,min=1"`
	Available      bool   `json:"available"`
	ImageId        string `json:"imageId"`
	ImageUrl       string `json:"imageUrl"`
}

type UpdateServiceRequest struct {
	SalonId        string  `json:"salonId" validate:"required,uuid4,min=1"`
	ProfessionalId string  `json:"professionalId" validate:"required,uuid4,min=1"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	Price          *int    `json:"price"`
	DurationInMin  *int    `json:"durationInMin"`
	Available      *bool   `json:"available"`
	ImageId        *string `json:"imageId"`
	ImageUrl       *string `json:"imageUrl"`
}

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

type ServiceRepository interface {
	FindManyByIds(ids []string) ([]*Service, error)
	FindById(id string) (*Service, error)
	FindBySalonId(salonId string) ([]*Service, error)
	Save(service *Service) (*Service, error)
}
