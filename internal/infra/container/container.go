package container

import (
	"time"

	"github.com/fio-de-navalha/fdn-back/config"
	"github.com/fio-de-navalha/fdn-back/internal/app"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database/repositories"
	"github.com/fio-de-navalha/fdn-back/internal/infra/providers/cloudflare"
)

func LoadSecurityQuestionService() *app.SecurityQuestionService {
	securityRepo := repositories.NewGormSecurityQuestionRepository()
	securityQuestionService := app.NewSecurityQuestionService(securityRepo)
	return securityQuestionService
}

func LoadVerificationCodeService() *app.VerificationCodeService {
	verificationCodeService := app.NewVerificationCodeService(2*time.Minute, 2*time.Minute)
	return verificationCodeService
}

func LoadCustomerService() *app.CustomerService {
	customerRepo := repositories.NewGormCustomerRepository()
	securityQuestionService := LoadSecurityQuestionService()
	verificationCodeService := LoadVerificationCodeService()
	customerService := app.NewCustomerService(customerRepo, *securityQuestionService, *verificationCodeService)
	return customerService
}

func LoadProfessionalService() *app.ProfessionalService {
	professionalRepo := repositories.NewGormProfessionalRepository()
	professionalService := app.NewProfessionalService(professionalRepo)
	return professionalService
}

func LoadSalonService() *app.SalonService {
	salonRepo := repositories.NewGormSalonRepository()
	salonMemberRepo := repositories.NewGormSalonMemberRepository()
	addressRepo := repositories.NewGormSalonAddressRepository()
	contactRepo := repositories.NewGormSalonContactRepository()
	periodRepo := repositories.NewGormSalonPeriodRepository()

	professionalService := LoadProfessionalService()

	salonService := app.NewSalonService(
		salonRepo,
		salonMemberRepo,
		addressRepo,
		contactRepo,
		periodRepo,
		*professionalService,
	)
	return salonService
}

func LoadServiceService() *app.ServiceService {
	serviceRepo := repositories.NewGormSalonServiceRepository()

	salonService := LoadSalonService()
	professionalService := LoadProfessionalService()

	cloudFlareService := cloudflare.NewCloudFlareService(
		config.CloudFlareBaseURL,
		config.CloudFlareAccountId,
		config.CloudFlareReadToken,
		config.CloudFlareEditToken,
	)

	serviceService := app.NewServiceService(
		serviceRepo,
		*salonService,
		*professionalService,
		cloudFlareService,
	)
	return serviceService
}

func LoadProductService() *app.ProductService {
	productRepo := repositories.NewGormSalonProductRepository()

	salonService := LoadSalonService()
	professionalService := LoadProfessionalService()

	cloudFlareService := cloudflare.NewCloudFlareService(
		config.CloudFlareBaseURL,
		config.CloudFlareAccountId,
		config.CloudFlareReadToken,
		config.CloudFlareEditToken,
	)

	productService := app.NewProductService(
		productRepo,
		*salonService,
		*professionalService,
		cloudFlareService,
	)
	return productService
}

func LoadAppointmentService() *app.AppointmentService {
	appointmentRepo := repositories.NewGormAppointmentRepository()

	professionalService := LoadProfessionalService()
	customerService := LoadCustomerService()
	salonService := LoadSalonService()
	serviceService := LoadServiceService()
	productService := LoadProductService()

	appointmentService := app.NewAppointmentService(
		appointmentRepo,
		*professionalService,
		*customerService,
		*serviceService,
		*productService,
		*salonService,
	)
	return appointmentService
}
