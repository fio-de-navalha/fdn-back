package barber

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/google/uuid"
)

type Barber struct {
	ID        string            `json:"id" gorm:"primaryKey"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	CreatedAt time.Time         `json:"createdAt"`
	Services  []service.Service `json:"services"`
	Products  []product.Product `json:"products"`
}

func NewBarber(input RegisterRequest) *Barber {
	return &Barber{
		ID:        uuid.NewString(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now(),
	}
}
