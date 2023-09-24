package professional

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/google/uuid"
)

type Professional struct {
	ID        string            `json:"id" gorm:"primaryKey"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	CreatedAt time.Time         `json:"createdAt"`
	Services  []service.Service `json:"services"`
}

func NewProfessional(input RegisterProfessionalRequest) *Professional {
	return &Professional{
		ID:        uuid.NewString(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now(),
	}
}
