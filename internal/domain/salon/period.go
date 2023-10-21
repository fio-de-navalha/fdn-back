package salon

import (
	"github.com/google/uuid"
)

type AddPeriodRequest struct {
	Day   int    `json:"day" validate:"required,min=0,max=6"`
	Open  string `json:"open" validate:"required,hourMinuteFormat"`
	Close string `json:"close" validate:"required,hourMinuteFormat"`
}

type UpdatePeriodRequest struct {
	Day   *int    `json:"day" validate:"required,min=0,max=6"`
	Open  *string `json:"open" validate:"required,hourMinuteFormat"`
	Close *string `json:"close" validate:"required,hourMinuteFormat"`
}

type Period struct {
	ID      string `json:"id"`
	SalonId string `json:"salonId"`
	Day     int    `json:"day"`
	Open    string `json:"open"`
	Close   string `json:"close"`

	Salon *Salon
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

type PeriodRepository interface {
	FindBySalonId(salonId string) ([]*Period, error)
	FindBySalonAndDay(salonId string, day int) (*Period, error)
	FindById(id string, salonId string) (*Period, error)
	Save(period *Period) (*Period, error)
	Delete(periodId string) error
}
