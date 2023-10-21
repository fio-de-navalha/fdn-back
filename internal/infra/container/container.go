package container

import (
	"github.com/fio-de-navalha/fdn-back/config"
	"github.com/fio-de-navalha/fdn-back/internal/app"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database/repositories"
	"github.com/fio-de-navalha/fdn-back/internal/infra/providers/cloudflare"
)

var (
	SalonService        *app.SalonService
	CustomerService     *app.CustomerService
	ProfessionalService *app.ProfessionalService
	ServiceService      *app.ServiceService
	ProductService      *app.ProductService
	AppointmentService  *app.AppointmentService
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

	CustomerService = app.NewCustomerService(customerRepo)
	ProfessionalService = app.NewProfessionalService(professionalRepo)
	SalonService = app.NewSalonService(
		salonRepo, salonMemberRepo, addressRepo, contactRepo, periodRepo, *ProfessionalService,
	)
	ServiceService = app.NewServiceService(
		serviceRepo, *SalonService, *ProfessionalService, cloudFlareService,
	)
	ProductService = app.NewProductService(
		productRepo, *SalonService, *ProfessionalService, cloudFlareService,
	)
	AppointmentService = app.NewAppointmentService(
		appointmentRepo,
		*ProfessionalService,
		*CustomerService,
		*ServiceService,
		*ProductService,
		*SalonService,
	)
}
