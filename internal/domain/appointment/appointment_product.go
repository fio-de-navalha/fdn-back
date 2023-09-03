package appointment

type AppointmentProduct struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	AppointmentID uint `json:"appointmentId"`
	ProductID     uint `json:"productId"`
}
