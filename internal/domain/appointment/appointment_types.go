package appointment

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
)

type AppointmentResponse struct {
	ID            uint              `json:"id" gorm:"primaryKey"`
	BarberId      string            `json:"barberId"`
	CustomerId    string            `json:"customerId"`
	DurationInMin int32             `json:"durationInMin"`
	StartsAt      time.Time         `json:"startsAt"`
	EndsAt        time.Time         `json:"endsAt"`
	CreatedAt     time.Time         `json:"createdAt"`
	Services      []service.Service `json:"services"`
	Products      []product.Product `json:"products"`
}
