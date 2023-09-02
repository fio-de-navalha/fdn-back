package container

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	gorm_repository "github.com/fio-de-navalha/fdn-back/internal/infra/database/gorm/repositories"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
)

var (
	CustomerService *application.CustomerService
	CustomerHandler *handlers.CustomerHandler

	BarberService *application.BarberService
	BarberHandler *handlers.BarberHandler

	ServiceService *application.ServiceService
	ServiceHandler *handlers.ServiceHandler

	ProductService *application.ProductService
	ProductHandler *handlers.ProductHandler
)

func LoadContainers() {
	customerRepo := gorm_repository.NewGormCustomerRepository()
	CustomerService = application.NewCustomerService(customerRepo)
	CustomerHandler = handlers.NewCustomerHandler(*CustomerService)

	barberRepo := gorm_repository.NewGormBarberRepository()
	BarberService = application.NewBarberService(barberRepo)
	BarberHandler = handlers.NewBarberHandler(*BarberService)

	serviceRepo := gorm_repository.NewGormServiceRepository()
	ServiceService = application.NewServiceService(serviceRepo, *BarberService)
	ServiceHandler = handlers.NewServiceHandler(*ServiceService)

	productRepo := gorm_repository.NewGormProductRepository()
	ProductService = application.NewProductService(productRepo, *BarberService)
	ProductHandler = handlers.NewProductHandler(*ProductService)
}
