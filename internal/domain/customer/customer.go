package customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type CustomerInput struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginInput struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string    `json:"access_token"`
	Customer    *Customer `json:"customer"`
}

func NewCustomer(input CustomerInput) *Customer {
	return &Customer{
		ID:        uuid.NewString(),
		Name:      input.Name,
		Phone:     input.Phone,
		Password:  input.Password,
		CreatedAt: time.Now(),
	}
}
