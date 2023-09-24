package container

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/config"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"github.com/fio-de-navalha/fdn-back/internal/infra/services/cloudflare"
)

var (
	CustomerService    *application.CustomerService
	BarberService      *application.BarberService
	ServiceService     *application.ServiceService
	ProductService     *application.ProductService
	AppointmentService *application.AppointmentService

	CloudFlareService *cloudflare.CloudFlareService
)

func LoadContainers() {
	customerRepo := database.NewGormCustomerRepository()
	addressRepo := database.NewGormAddressRepository()
	contactRepo := database.NewGormContactRepository()
	barberRepo := database.NewGormBarberRepository()
	serviceRepo := database.NewGormServiceRepository()
	productRepo := database.NewGormProductRepository()
	appointmentRepo := database.NewGormAppointmentRepository()

	cloudFlareService := cloudflare.NewCloudFlareService(
		config.CloudFlareBaseURL,
		config.CloudFlareAccountId,
		config.CloudFlareReadToken,
		config.CloudFlareEditToken,
	)

	CustomerService = application.NewCustomerService(customerRepo)
	BarberService = application.NewBarberService(barberRepo, addressRepo, contactRepo)
	ServiceService = application.NewServiceService(serviceRepo, *BarberService, cloudFlareService)
	ProductService = application.NewProductService(productRepo, *BarberService)
	AppointmentService = application.NewAppointmentService(
		appointmentRepo,
		*BarberService,
		*CustomerService,
		*ServiceService,
		*ProductService,
	)
}
