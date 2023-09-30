package salon

type PeriodRepository interface {
	FindBySalonId(salonId string) ([]*Period, error)
	FindById(id string, salonId string) (*Period, error)
	Save(period *Period) (*Period, error)
	Delete(periodId string) error
}
