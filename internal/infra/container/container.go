package container

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	gorm_repository "github.com/fio-de-navalha/fdn-back/internal/infra/database/gorm/repositories"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
)

var (
	CustomerService *application.CustomerServices
	CustomerHandler *handlers.CustomerHandler

	BarberService *application.BarberService
	BarberHandler *handlers.BarberHandler
)

func LoadContainers() {
	customerRepo := gorm_repository.NewGormCustomerRepository()
	CustomerService = application.NewCustomerServices(customerRepo)
	CustomerHandler = handlers.NewCustomerHandler(*CustomerService)

	barberRepo := gorm_repository.NewGormBarberRepository()
	BarberService = application.NewBarberService(barberRepo)
	BarberHandler = handlers.NewBarberHandler(*BarberService)
}
