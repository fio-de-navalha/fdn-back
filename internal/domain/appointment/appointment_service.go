package appointment

type AppointmentService struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	AppointmentID uint `json:"appointmentId"`
	ServiceID     uint `json:"serviceId"`
}
