package application

import (
	"fmt"

	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
	"github.com/fio-de-navalha/fdn-back/internal/domain/barber"
)

type AppointmentService struct {
	appointmentRepository appointment.AppointmentRepository
	// barberService         BarberService
	// customerService       CustomerService
	// serviceService        ServiceService
	// productService        ProductService
}

func NewAppointmentService(appointmentRepository appointment.AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		appointmentRepository: appointmentRepository,
	}
}

func (s *AppointmentService) GetBarberAppointment(barberId string) ([]*appointment.Appointment, error) {
	a, err := s.appointmentRepository.FindByBarberId(barberId)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return a, nil
}

func (s *AppointmentService) GetCustomerAppointment(customerId string) ([]*appointment.Appointment, error) {
	a, err := s.appointmentRepository.FindByCustomerId(customerId)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return a, nil
}

func (s *AppointmentService) GetAppointment(id uint) (*appointment.Appointment, error) {
	a, err := s.appointmentRepository.FindById(id)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return a, nil
}

func (s *AppointmentService) CreateApppointment(input barber.RegisterRequest) error {
	// Check barber exists
	// Check customer exists
	// Check services exists
	// Check products exists
	// Check if date time is available
	// Create appointment

	// barberExists, err := s.appointmentRepository.Save()
	// if err != nil {
	// 	return err
	// }
	// if barberExists != nil {
	// 	return errors.New("barber alredy exists")
	// }

	// hashedPassword, err := cryptography.HashPassword(input.Password)
	// if err != nil {
	// 	return err
	// }

	// input = barber.RegisterRequest{
	// 	Name:     input.Name,
	// 	Email:    input.Email,
	// 	Password: hashedPassword,
	// }

	// bar := barber.NewBarber(input)
	// _, err = s.appointmentRepository.Save(bar)
	// if err != nil {
	// 	// TODO: add better error handling
	// 	fmt.Println(err)
	// }
	return nil
}
