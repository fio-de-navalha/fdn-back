package appointment

type AppointmentService struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	AppointmentID string `json:"appointmentId"`
	ServiceID     string `json:"serviceId"`
}

func NewAppointmentService(appointmentId string, serviceId string) *AppointmentService {
	return &AppointmentService{
		AppointmentID: appointmentId,
		ServiceID:     serviceId,
	}
}
