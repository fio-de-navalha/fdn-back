package container

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
)

var (
	CustomerService    *application.CustomerService
	BarberService      *application.BarberService
	ServiceService     *application.ServiceService
	ProductService     *application.ProductService
	AppointmentService *application.AppointmentService
)

func LoadContainers() {
	customerRepo := database.NewGormCustomerRepository()
	addressRepo := database.NewGormAddressRepository()
	contactRepo := database.NewGormContactRepository()
	barberRepo := database.NewGormBarberRepository()
	serviceRepo := database.NewGormServiceRepository()
	productRepo := database.NewGormProductRepository()
	appointmentRepo := database.NewGormAppointmentRepository()

	CustomerService = application.NewCustomerService(customerRepo)
	BarberService = application.NewBarberService(barberRepo, addressRepo, contactRepo)
	ServiceService = application.NewServiceService(serviceRepo, *BarberService)
	ProductService = application.NewProductService(productRepo, *BarberService)
	AppointmentService = application.NewAppointmentService(
		appointmentRepo,
		*BarberService,
		*CustomerService,
		*ServiceService,
		*ProductService,
	)
}
