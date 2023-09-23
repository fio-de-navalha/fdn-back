package customer

import "time"

type CustomerResponse struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=30"`
	Phone    string `json:"phone" validate:"required,min=9,max=15"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required,min=9,max=15"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthCustomerResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
}

type AuthResponse struct {
	AccessToken string               `json:"access_token"`
	Customer    AuthCustomerResponse `json:"customer"`
}
