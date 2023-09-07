package application

import (
	"errors"
	"fmt"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
)

type AppointmentService struct {
	appointmentRepository appointment.AppointmentRepository
	barberService         BarberService
	customerService       CustomerService
	serviceService        ServiceService
	productService        ProductService
}

func NewAppointmentService(
	appointmentRepository appointment.AppointmentRepository,
	barberService BarberService,
	customerService CustomerService,
	serviceService ServiceService,
	productService ProductService,
) *AppointmentService {
	return &AppointmentService{
		appointmentRepository: appointmentRepository,
		barberService:         barberService,
		customerService:       customerService,
		serviceService:        serviceService,
		productService:        productService,
	}
}

func (s *AppointmentService) GetBarberAppointments(barberId string) ([]*appointment.Appointment, error) {
	fmt.Println(barberId)
	a, err := s.appointmentRepository.FindByBarberId(barberId)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return a, nil
}

func (s *AppointmentService) GetCustomerAppointments(customerId string) ([]*appointment.Appointment, error) {
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

func (s *AppointmentService) CreateApppointment(input appointment.CreateAppointmentRequest) error {
	// Check barber exists
	_, err := s.barberService.GetBarberById(input.BarberId)
	if err != nil {
		return errors.New("barber not found")
	}

	// Check customer exists
	_, err = s.customerService.GetCustomerById(input.CustomerId)
	if err != nil {
		return errors.New("customer not found")
	}

	// Check services exists and is availeble
	var durationInMin int32 // Change durationInMin type to int
	var servicesToSave []uint
	services, err := s.serviceService.getManyServices(input.ServiceIds)
	if err != nil {
		return err
	}
	for _, v := range services {
		if v.Available {
			durationInMin += v.DurationInMin
			servicesToSave = append(servicesToSave, v.ID)
		}
	}

	// Check products exists
	var productsToSave []uint
	products, err := s.productService.getManyProducts(input.ProductIds)
	if err != nil {
		return err
	}
	for _, v := range products {
		if v.Available {
			productsToSave = append(productsToSave, v.ID)
		}
	}

	// Check if date time is available
	endsAt := input.StartsAt.Add(time.Minute * time.Duration(durationInMin))
	appos, err := s.appointmentRepository.FindByDates(input.StartsAt, endsAt)
	if err != nil {
		return err
	}
	if len(appos) > 0 {
		return errors.New("time box not available")
	}

	// Create appointment
	appo := appointment.NewAppointment(
		input.BarberId,
		input.CustomerId,
		int(durationInMin),
		input.StartsAt,
		endsAt,
	)
	_, err = s.appointmentRepository.Save(appo)
	if err != nil {
		return err
	}

	return nil
}
