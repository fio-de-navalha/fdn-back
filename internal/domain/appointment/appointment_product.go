package appointment

type AppointmentProduct struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	AppointmentID string `json:"appointmentId"`
	ProductID     string `json:"productId"`
}

func NewAppointmentProduct(appointmentId string, productId string) *AppointmentProduct {
	return &AppointmentProduct{
		AppointmentID: appointmentId,
		ProductID:     productId,
	}
}
