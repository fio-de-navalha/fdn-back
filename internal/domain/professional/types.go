package professional

import (
	"time"
)

type ProfessionalResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterProfessionalRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=30"`
	Email    string `json:"email" validate:"required,email,min=3,max=30"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginProfessionalRequest struct {
	Email    string `json:"email" validate:"required,email,min=3,max=30"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthProfessionalResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type AuthResponse struct {
	AccessToken  string                   `json:"access_token"`
	Professional AuthProfessionalResponse `json:"professional"`
}
