package application

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fio-de-navalha/fdn-back/config"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"golang.org/x/exp/slices"
)

type AppointmentService struct {
	appointmentRepository appointment.AppointmentRepository
	professionalService   ProfessionalService
	customerService       CustomerService
	serviceService        ServiceService
	productService        ProductService
	salonService          SalonService
}

func NewAppointmentService(
	appointmentRepository appointment.AppointmentRepository,
	professionalService ProfessionalService,
	customerService CustomerService,
	serviceService ServiceService,
	productService ProductService,
	salonService SalonService,
) *AppointmentService {
	return &AppointmentService{
		appointmentRepository: appointmentRepository,
		professionalService:   professionalService,
		customerService:       customerService,
		serviceService:        serviceService,
		productService:        productService,
		salonService:          salonService,
	}
}

func (s *AppointmentService) GetProfessionalAppointments(professionalId string, startsAt time.Time) ([]*appointment.Appointment, error) {
	log.Println("[AppointmentService.GetProfessionalAppointments] - Getting appointments from professional:", professionalId)
	endsAt := time.Date(
		startsAt.Year(),
		startsAt.Month(),
		startsAt.Day(),
		constants.EndsAtHour, constants.EndsAtMinute, 0, 0,
		startsAt.Location(),
	)
	a, err := s.appointmentRepository.FindByProfessionalId(professionalId, startsAt, endsAt)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AppointmentService) GetCustomerAppointments(customerId string) ([]*appointment.Appointment, error) {
	log.Println("[AppointmentService.GetCustomerAppointments] - Getting appointments from customer:", customerId)
	a, err := s.appointmentRepository.FindByCustomerId(customerId)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AppointmentService) GetAppointment(id string) (*appointment.Appointment, error) {
	log.Println("[AppointmentService.GetAppointment] - Getting appointment:", id)
	a, err := s.appointmentRepository.FindByIdWithJoins(id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AppointmentService) CreateApppointment(input appointment.CreateAppointmentRequest) error {
	type chanResultService struct {
		IDs         []string
		Duration    int
		TotalAmount int
	}
	type chanResultProduct struct {
		IDs []string
	}

	resultServiceChan := make(chan chanResultService, 1)
	resultProductChan := make(chan chanResultProduct, 1)
	errs := make(chan error, 7)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		if err := s.validateEntity("professional", input.ProfessionalId, func(id string) (interface{}, error) {
			return s.professionalService.GetProfessionalById(id)
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

	go func() {
		defer wg.Done()
		log.Println("[AppointmentService.CreateApppointment] - Validating services:", input.ServiceIds)
		services, err := s.serviceService.getManyServices(input.ServiceIds)
		if err != nil {
			errs <- err
		}
		if len(services) == 0 {
			errs <- &utils.AppError{
				Code:    constants.SERVICE_NOT_FOUND_ERROR_CODE,
				Message: constants.SERVICE_NOT_FOUND_ERROR_MESSAGE,
			}
		}
		idsToSave, duration, totalAmount := s.serviceService.ValidateServicesAvailability(services)
		if err := s.validateAssociation("services", input.ServiceIds, idsToSave); err != nil {
			errs <- err
		}
		resultServiceChan <- chanResultService{
			IDs:         idsToSave,
			Duration:    duration,
			TotalAmount: totalAmount,
		}
	}()

	go func() {
		defer wg.Done()
		log.Println("[AppointmentService.CreateApppointment] - Validating products:", input.ProductIds)
		products, err := s.productService.getManyProducts(input.ProductIds)
		if err != nil {
			errs <- err
		}
		idsToSave := s.productService.validateProductsAvailability(products)
		if err := s.validateAssociation("products", input.ProductIds, idsToSave); err != nil {
			errs <- err
		}
		resultProductChan <- chanResultProduct{
			IDs: idsToSave,
		}
	}()

	wg.Wait()
	close(errs)
	close(resultServiceChan)
	close(resultProductChan)
	for err := range errs {
		log.Println("[AppointmentService.CreateApppointment] - Validation error:", err)
	}

	for err := range errs {
		return err
	}

	resultService := <-resultServiceChan
	resultProduct := <-resultProductChan

	endsAt := input.StartsAt.Add(time.Minute * time.Duration(resultService.Duration))
	log.Println("[AppointmentService.CreateApppointment] - Validating appointment time range availability for:")
	log.Println("[AppointmentService.CreateApppointment] - StartsAt:", input.StartsAt)
	log.Println("[AppointmentService.CreateApppointment] - EndsAt:", endsAt)
	if err := s.validateAppointmentTimeRange(input.SalonId, input.StartsAt, endsAt); err != nil {
		return err
	}

	log.Println("[AppointmentService.CreateApppointment] - Creating appointment")
	appo := appointment.NewAppointment(
		input.ProfessionalId,
		input.CustomerId,
		resultService.Duration,
		resultService.TotalAmount,
		input.StartsAt,
		endsAt,
	)
	servicesToSave, productsToSave := s.newAppointmentItems(appo.ID, resultService.IDs, resultProduct.IDs)
	if _, err := s.appointmentRepository.Save(appo, servicesToSave, productsToSave); err != nil {
		return err
	}

	return nil
}

func (s *AppointmentService) CancelApppointment(requesterId string, appointmentId string) error {
	log.Println("[AppointmentService.CancelApppointment] - Validating appointment:", appointmentId)
	appo, err := s.appointmentRepository.FindById(appointmentId)
	if err != nil {
		return err
	}
	if appo == nil {
		return &utils.AppError{
			Code:    constants.APPOINTMENT_NOT_FOUND_ERROR_CODE,
			Message: constants.APPOINTMENT_NOT_FOUND_ERROR_MESSAGE,
		}
	}

	if requesterId != appo.ProfessionalId && requesterId != appo.CustomerId {
		return &utils.AppError{
			Code:    constants.PERMISSION_DENIED_ERROR_CODE,
			Message: constants.PERMISSION_DENIED_ERROR_MESSAGE,
		}
	}
	if appo.StartsAt.Before(time.Now()) {
		return &utils.AppError{
			Code:    constants.CANNOT_CANCEL_PAST_APPOINTMENT_ERROR_CODE,
			Message: constants.CANNOT_CANCEL_PAST_APPOINTMENT_ERROR_MESSAGE,
		}
	}

	log.Println("[AppointmentService.CancelApppointment] - Canceling appointment:", appointmentId)
	if _, err := s.appointmentRepository.Cancel(appo); err != nil {
		return err
	}

	return nil
}

func (s *AppointmentService) validateEntity(
	context string,
	param string,
	fn func(string) (interface{}, error),
) error {
	_, err := fn(param)
	if err != nil {
		var code int
		if context == "professional" {
			code = constants.PROFESSIONAL_NOT_FOUND_ERROR_CODE
		}
		if context == "customer" {
			code = constants.CUSTOMER_NOT_FOUND_ERROR_CODE
		}
		return &utils.AppError{
			Code:    code,
			Message: context + " not found",
		}
	}
	return nil
}

func (s *AppointmentService) validateAppointmentTimeRange(salonId string, startsAt, endsAt time.Time) error {
	day := startsAt.Weekday()
	log.Println("day string:", day)
	log.Println("day int:", int(day))
	per, err := s.salonService.GetSalonPeriodByDay(salonId, int(day))
	if err != nil {
		return err
	}
	log.Println("period:", per)
	if per == nil {
		log.Println("[AppointmentService.validateAppointmentTimeRange] - per == nil")
		return &utils.AppError{
			Code:    constants.SALON_CLOSED_ERROR_CODE,
			Message: constants.SALON_CLOSED_ERROR_MESSAGE,
		}
	}

	openHourStr := strings.Split(per.Open, ":")
	openHour, _ := strconv.Atoi(openHourStr[0])
	openMinute, _ := strconv.Atoi(openHourStr[1])
	openTime := time.Date(startsAt.Year(), startsAt.Month(), startsAt.Day(), openHour, openMinute, 0, 0, config.GMTMinus3)
	if startsAt.Before(openTime) {
		log.Println("[AppointmentService.validateAppointmentTimeRange] - startsAt.Before(openTime)")
		return &utils.AppError{
			Code:    constants.SALON_CLOSED_ERROR_CODE,
			Message: constants.SALON_CLOSED_ERROR_MESSAGE,
		}
	}

	closeHourStr := strings.Split(per.Close, ":")
	closeHour, _ := strconv.Atoi(closeHourStr[0])
	closeMinute, _ := strconv.Atoi(closeHourStr[1])
	closeTime := time.Date(startsAt.Year(), startsAt.Month(), startsAt.Day(), closeHour, closeMinute, 0, 0, config.GMTMinus3)
	log.Println("closeTime default", closeTime)
	log.Println("closeTime convert", utils.ConvertToGMTMinus3(closeTime))
	if startsAt.After(closeTime) {
		log.Println("[AppointmentService.validateAppointmentTimeRange] - startsAt.After(closeTime)")
		return &utils.AppError{
			Code:    constants.SALON_CLOSED_ERROR_CODE,
			Message: constants.SALON_CLOSED_ERROR_MESSAGE,
		}
	}

	appos, err := s.appointmentRepository.FindByDates(startsAt, endsAt)
	if err != nil {
		return err
	}
	if len(appos) > 0 {
		return &utils.AppError{
			Code:    constants.APPOINTMENT_TIME_UNAVAILABLE_ERROR_CODE,
			Message: constants.APPOINTMENT_TIME_UNAVAILABLE_ERROR_MESSAGE,
		}
	}
	return nil
}

func (s *AppointmentService) validateAssociation(context string, input []string, idsToSave []string) error {
	var itemNotFound []string
	for _, id := range input {
		if !slices.Contains(idsToSave, id) {
			itemNotFound = append(itemNotFound, id)
			continue
		}
	}
	if len(itemNotFound) > 0 {
		var code int
		if context == "services" {
			code = constants.SERVICE_NOT_FOUND_ERROR_CODE
		}
		if context == "products" {
			code = constants.PRODUCT_NOT_FOUND_ERROR_CODE
		}
		return &utils.AppError{
			Code:    code,
			Message: context + " not found:" + strings.Join(itemNotFound, ", "),
		}
	}
	return nil
}

func (s *AppointmentService) newAppointmentItems(
	appoId string,
	servicesIds []string,
	productsIds []string,
) (
	[]*appointment.AppointmentService,
	[]*appointment.AppointmentProduct,
) {
	var servicesToSave []*appointment.AppointmentService
	var productsToSave []*appointment.AppointmentProduct

	for _, v := range servicesIds {
		ser := appointment.NewAppointmentService(appoId, v)
		servicesToSave = append(servicesToSave, ser)
	}
	for _, v := range productsIds {
		pro := appointment.NewAppointmentProduct(appoId, v)
		productsToSave = append(productsToSave, pro)
	}

	return servicesToSave, productsToSave
}
