package container

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/config"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"github.com/fio-de-navalha/fdn-back/internal/infra/providers/cloudflare"
)

var (
	SalonService        *application.SalonService
	CustomerService     *application.CustomerService
	ProfessionalService *application.ProfessionalService
	ServiceService      *application.ServiceService
	ProductService      *application.ProductService
	AppointmentService  *application.AppointmentService
)

func LoadContainers() {
	salonRepo := database.NewGormSalonRepository()
	salonMemberRepo := database.NewGormSalonMemberRepository()
	customerRepo := database.NewGormCustomerRepository()
	addressRepo := database.NewGormAddressRepository()
	contactRepo := database.NewGormContactRepository()
	periodRepo := database.NewGormPeriodRepository()
	professionalRepo := database.NewGormProfessionalRepository()
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
	ProfessionalService = application.NewProfessionalService(professionalRepo)
	SalonService = application.NewSalonService(
		salonRepo, salonMemberRepo, addressRepo, contactRepo, periodRepo, *ProfessionalService,
	)
	ServiceService = application.NewServiceService(
		serviceRepo, *SalonService, *ProfessionalService, cloudFlareService,
	)
	ProductService = application.NewProductService(
		productRepo, *SalonService, *ProfessionalService, cloudFlareService,
	)
	AppointmentService = application.NewAppointmentService(
		appointmentRepo,
		*ProfessionalService,
		*CustomerService,
		*ServiceService,
		*ProductService,
		*SalonService,
	)
}
