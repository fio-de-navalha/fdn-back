package container

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	gorm_repository "github.com/fio-de-navalha/fdn-back/internal/infra/database/gorm/repositories"
)

var (
	CustomerService    *application.CustomerService
	BarberService      *application.BarberService
	ServiceService     *application.ServiceService
	ProductService     *application.ProductService
	AppointmentService *application.AppointmentService
)

func LoadContainers() {
	customerRepo := gorm_repository.NewGormCustomerRepository()
	CustomerService = application.NewCustomerService(customerRepo)

	addressRepo := gorm_repository.NewGormAddressRepository()
	contactRepo := gorm_repository.NewGormContactRepository()
	barberRepo := gorm_repository.NewGormBarberRepository()
	BarberService = application.NewBarberService(barberRepo, addressRepo, contactRepo)

	serviceRepo := gorm_repository.NewGormServiceRepository()
	ServiceService = application.NewServiceService(serviceRepo, *BarberService)

	productRepo := gorm_repository.NewGormProductRepository()
	ProductService = application.NewProductService(productRepo, *BarberService)

	appointmentRepo := gorm_repository.NewGormAppointmentRepository()
	AppointmentService = application.NewAppointmentService(
		appointmentRepo,
		*BarberService,
		*CustomerService,
		*ServiceService,
		*ProductService,
	)

}
