package container

import (
	"github.com/fio-de-navalha/fdn-back/config"
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database/repositories"
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
	salonRepo := repositories.NewGormSalonRepository()
	salonMemberRepo := repositories.NewGormSalonMemberRepository()
	customerRepo := repositories.NewGormCustomerRepository()
	addressRepo := repositories.NewGormSalonAddressRepository()
	contactRepo := repositories.NewGormSalonContactRepository()
	periodRepo := repositories.NewGormSalonPeriodRepository()
	professionalRepo := repositories.NewGormProfessionalRepository()
	serviceRepo := repositories.NewGormSalonServiceRepository()
	productRepo := repositories.NewGormSalonProductRepository()
	appointmentRepo := repositories.NewGormAppointmentRepository()

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
