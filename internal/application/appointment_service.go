package application

import (
	"errors"
	"log"
	"strings"
	"sync"
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

	log.Println("[application.GetBarberAppointments] - Getting appointments from barber:", barberId)
	a, err := s.appointmentRepository.FindByBarberId(barberId, startsAt, endsAt)
	if err != nil {
		log.Println("[application.GetBarberAppointments] - Error when getting appointments from barber:", barberId)
		return nil, err
	}
	log.Println("[application.GetBarberAppointments] - Successfully got appointments from barber:", barberId)
	return a, nil
}

func (s *AppointmentService) GetCustomerAppointments(customerId string) ([]*appointment.Appointment, error) {
	log.Println("[application.GetCustomerAppointments] - Getting appointments from customer:", customerId)
	a, err := s.appointmentRepository.FindByCustomerId(customerId)
	if err != nil {
		log.Println("[application.GetCustomerAppointments] - Error when getting appointments from customer:", customerId)
		return nil, err
	}
	log.Println("[application.GetCustomerAppointments] - Successfully got appointments from customer:", customerId)
	return a, nil
}

func (s *AppointmentService) GetAppointment(id string) (*appointment.Appointment, error) {
	log.Println("[application.GetAppointment] - Getting appointment:", id)
	a, err := s.appointmentRepository.FindById(id)
	if err != nil {
		log.Println("[application.GetAppointment] - Error when getting appointment:", id)
		return nil, err
	}
	log.Println("[application.GetAppointment] - Successfully got appointment:", id)
	return a, nil
}

func (s *AppointmentService) CreateApppointment(input appointment.CreateAppointmentRequest) error {
	var wg sync.WaitGroup
	errs := make(chan error, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := s.validateEntity("barber", input.BarberId, func(id string) (interface{}, error) {
			return s.barberService.GetBarberById(id)
		}); err != nil {
			errs <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err := s.validateEntity("customer", input.CustomerId, func(id string) (interface{}, error) {
			return s.customerService.GetCustomerById(id)
		}); err != nil {
			errs <- err
		}
	}()

	wg.Wait()
	close(errs)

	for err := range errs {
		return err
	}

	log.Println("[application.CreateApppointment] - Validating services:", input.ServiceIds)
	services, err := s.serviceService.getManyServices(input.ServiceIds)
	if err != nil {
		return err
	}
	if len(services) == 0 {
		return errors.New("services not found")
	}
	servicesIdsToSave, durationInMin := s.serviceService.ValidateServicesAvailability(services)
	if err := validateAssociation("services", input.ServiceIds, servicesIdsToSave); err != nil {
		return err
	}

	log.Println("[application.CreateApppointment] - Validating products:", input.ProductIds)
	products, err := s.productService.getManyProducts(input.ProductIds)
	if err != nil {
		return err
	}
	productsIdsToSave := s.productService.ValidateProductsAvailability(products)
	if err := validateAssociation("products", input.ProductIds, productsIdsToSave); err != nil {
		return err
	}

	log.Println("[application.CreateApppointment] - Validating appointment time range availability")
	endsAt := input.StartsAt.Add(time.Minute * time.Duration(durationInMin))
	if err := s.validateAppointmentTimeRange(input.StartsAt, endsAt); err != nil {
		return err
	}

	log.Println("[application.CreateApppointment] - Creating appointment")
	appo := appointment.NewAppointment(
		input.BarberId,
		input.CustomerId,
		durationInMin,
		input.StartsAt,
		endsAt,
	)
	var servicesToSave []*appointment.AppointmentService
	var productsToSave []*appointment.AppointmentProduct
	for _, v := range servicesIdsToSave {
		ser := appointment.NewAppointmentService(appo.ID, v)
		servicesToSave = append(servicesToSave, ser)
	}
	for _, v := range productsIdsToSave {
		pro := appointment.NewAppointmentProduct(appo.ID, v)
		productsToSave = append(productsToSave, pro)
	}
	if _, err := s.appointmentRepository.Save(appo, servicesToSave, productsToSave); err != nil {
		return err
	}

	log.Println("[application.CreateApppointment] - Successfully created appointment")
	return nil
}

func (s *AppointmentService) validateEntity(
	context string,
	param string,
	fn func(string) (interface{}, error),
) error {
	log.Println("[application.CreateApppointment] - Validating", context, ":", param)
	_, err := fn(param)
	if err != nil {
		return errors.New(context + " not found")
	}

	return nil
}

func (s *AppointmentService) validateAppointmentTimeRange(startsAt, endsAt time.Time) error {
	appos, err := s.appointmentRepository.FindByDates(startsAt, endsAt)
	if err != nil {
		return err
	}
	if len(appos) > 0 {
		return errors.New("time box not available")
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
