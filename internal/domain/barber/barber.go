package barber

import (
	"time"

	"github.com/google/uuid"
)

type Barber struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type (
	BarberInput struct {
		Name     string `json:"name" validate:"required,min=8,max=32"`
		Email    string `json:"email" validate:"required,min=8,max=32"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginInput struct {
		Email    string `json:"email" validate:"required,min=8,max=32"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginResponse struct {
		AccessToken string  `json:"access_token"`
		Barber      *Barber `json:"barber"`
	}
)

func NewBarber(input BarberInput) *Barber {
	return &Barber{
		ID:        uuid.NewString(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now(),
	}
}
