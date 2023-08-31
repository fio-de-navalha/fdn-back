package container

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	gorm_repository "github.com/fio-de-navalha/fdn-back/internal/infra/database/gorm/repositories"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
)

var (
	CustomerService *application.CustomerServices
	CustomerHandler *handlers.CustomerHandler
	AuthHandler     *handlers.AuthHandler
)

func LoadContainers() {
	customerRepo := gorm_repository.NewGormCustomerRepository()
	CustomerService = application.NewCustomerServices(customerRepo)
	CustomerHandler = handlers.NewCustomerHandler(*CustomerService)
	AuthHandler = handlers.NewAuthHandler(*CustomerService)
}
