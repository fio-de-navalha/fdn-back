package barber

type BarberInput struct {
	Name     string `json:"name" validate:"required,min=8,max=32"`
	Email    string `json:"email" validate:"required,min=8,max=32"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,min=8,max=32"`
	Password string `json:"password" validate:"required,min=6""`
}

type LoginResponse struct {
	AccessToken string  `json:"access_token"`
	Barber      *Barber `json:"barber"`
}
