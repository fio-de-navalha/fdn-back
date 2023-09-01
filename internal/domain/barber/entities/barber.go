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

type BarberInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string  `json:"access_token"`
	Barber      *Barber `json:"barber"`
}

func NewBarber(input BarberInput) *Barber {
	return &Barber{
		ID:       uuid.NewString(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}
