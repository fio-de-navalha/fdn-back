package config

import (
	migrations "github.com/fio-de-navalha/fdn-back/db"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
)

func LoadConfigs() {
	loadEnvVariables()
	database.Connect()
	migrations.Migrate()
	container.LoadContainers()
}
