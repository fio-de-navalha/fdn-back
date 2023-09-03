package customer

type CustomerInput struct {
	Name     string `json:"name" validate:"required,min=3,max=30"`
	Phone    string `json:"phone" validate:"required,min=9,max=15"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginInput struct {
	Phone    string `json:"phone" validate:"required,min=9,max=15"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	AccessToken string    `json:"access_token"`
	Customer    *Customer `json:"customer"`
}
