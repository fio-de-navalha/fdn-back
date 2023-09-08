package application

import (
	"errors"
	"strings"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
	"golang.org/x/exp/slices"
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

func (s *AppointmentService) GetBarberAppointments(barberId string, startsAt time.Time) ([]*appointment.Appointment, error) {
	endsAt := time.Date(
		startsAt.Year(),
		startsAt.Month(),
		startsAt.Day(),
		constants.EndsAtHour, constants.EndsAtMinute, 0, 0,
		startsAt.Location(),
	)

	a, err := s.appointmentRepository.FindByBarberId(barberId, startsAt, endsAt)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AppointmentService) GetCustomerAppointments(customerId string) ([]*appointment.Appointment, error) {
	a, err := s.appointmentRepository.FindByCustomerId(customerId)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AppointmentService) GetAppointment(id string) (*appointment.Appointment, error) {
	a, err := s.appointmentRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AppointmentService) CreateApppointment(input appointment.CreateAppointmentRequest) error {
	_, err := s.barberService.GetBarberById(input.BarberId)
	if err != nil {
		return errors.New("barber not found")
	}

	_, err = s.customerService.GetCustomerById(input.CustomerId)
	if err != nil {
		return errors.New("customer not found")
	}

	var durationInMin int
	var servicesIdsToSave []string
	services, err := s.serviceService.getManyServices(input.ServiceIds)
	if err != nil {
		return err
	}
	if len(services) == 0 {
		return errors.New("services not found")
	}
	for _, v := range services {
		if v.Available {
			durationInMin += v.DurationInMin
			servicesIdsToSave = append(servicesIdsToSave, v.ID)
		}
	}
	err = validateAssociation("services", input.ServiceIds, servicesIdsToSave)
	if err != nil {
		return err
	}

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
	err = validateAssociation("products", input.ProductIds, productsIdsToSave)
	if err != nil {
		return err
	}

	endsAt := input.StartsAt.Add(time.Minute * time.Duration(durationInMin))
	appos, err := s.appointmentRepository.FindByDates(input.StartsAt, endsAt)
	if err != nil {
		return err
	}
	if len(appos) > 0 {
		return errors.New("time box not available")
	}

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

func validateAssociation(context string, input []string, idsToSave []string) error {
	var itemNotFound []string
	for _, id := range input {
		if !slices.Contains(idsToSave, id) {
			itemNotFound = append(itemNotFound, id)
			continue
		}
	}
	if len(itemNotFound) > 0 {
		return errors.New(context + " not found:" + strings.Join(itemNotFound, ", "))
	}
	return nil
}
