package config

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database/migrations"
)

func LoadConfigs() {
	loadEnvVariables()
	database.Connect()
	migrations.Migrate()
	container.LoadContainers()
}
