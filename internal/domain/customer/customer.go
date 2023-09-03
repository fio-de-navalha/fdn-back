package customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
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
