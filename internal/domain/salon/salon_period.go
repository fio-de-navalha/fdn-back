package salon

import (
	"time"

	"github.com/google/uuid"
)

type Period struct {
	ID      string    `json:"id"`
	SalonId string    `json:"salonId"`
	Day     int       `json:"day"`
	Open    time.Time `json:"open"`
	Close   time.Time `json:"close"`
}

func NewPeriod(salonId string, day int, open time.Time, close time.Time) *Period {
	return &Period{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Day:     day,
		Open:    open,
		Close:   close,
	}
}
