package config

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
)

func LoadConfigs() {
	loadEnvVariables()
	database.Connect()
	container.LoadContainers()
}
