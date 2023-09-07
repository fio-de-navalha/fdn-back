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

func (s *AppointmentService) GetAppointment(id string) (*appointment.Appointment, error) {
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
	var durationInMin int // Change durationInMin type to int
	var servicesIdsToSave []string
	services, err := s.serviceService.getManyServices(input.ServiceIds)
	if err != nil {
		return err
	}
	for _, v := range services {
		if v.Available {
			durationInMin += v.DurationInMin
			servicesIdsToSave = append(servicesIdsToSave, v.ID)
		}
	}

	// Check products exists
	var productsIdsToSave []string
	products, err := s.productService.getManyProducts(input.ProductIds)
	if err != nil {
		return err
	}
	for _, v := range products {
		if v.Available {
			productsIdsToSave = append(productsIdsToSave, v.ID)
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
		durationInMin,
		input.StartsAt,
		endsAt,
	)

	var servicesToSave []*appointment.AppointmentService
	for _, v := range servicesIdsToSave {
		ser := appointment.NewAppointmentService(appo.ID, v)
		servicesToSave = append(servicesToSave, ser)
	}

	var productsToSave []*appointment.AppointmentProduct
	for _, v := range productsIdsToSave {
		pro := appointment.NewAppointmentProduct(appo.ID, v)
		productsToSave = append(productsToSave, pro)
	}

	_, err = s.appointmentRepository.Save(appo, servicesToSave, productsToSave)
	if err != nil {
		return err
	}

	return nil
}
