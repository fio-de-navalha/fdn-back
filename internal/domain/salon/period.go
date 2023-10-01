package salon

import (
	"github.com/google/uuid"
)

type AddPeriodRequest struct {
	Day   int    `json:"day" validate:"required,min=1,max=7"`
	Open  string `json:"open" validate:"required,hourMinuteFormat"`
	Close string `json:"close" validate:"required,hourMinuteFormat"`
}

type UpdatePeriodRequest struct {
	Day   *int    `json:"day" validate:"required,min=1,max=7"`
	Open  *string `json:"open" validate:"required,hourMinuteFormat"`
	Close *string `json:"close" validate:"required,hourMinuteFormat"`
}

type Period struct {
	ID      string `json:"id"`
	SalonId string `json:"salonId"`
	Day     int    `json:"day"`
	Open    string `json:"open"`
	Close   string `json:"close"`
}

func NewPeriod(salonId string, day int, open string, close string) *Period {
	return &Period{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Day:     day,
		Open:    open,
		Close:   close,
	}
}
