package appointment

import (
	"time"
)

type Appointment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	BarberId      string    `json:"barberId"`
	CustomerId    string    `json:"customerId"`
	DurationInMin int32     `json:"durationInMin"`
	StartsAt      time.Time `json:"startsAt"`
	EndsAt        time.Time `json:"endsAt"`
	CreatedAt     time.Time `json:"createdAt"`
}
