package config

import (
	"github.com/fio-de-navalha/fdn-back/infra/database"
	"github.com/fio-de-navalha/fdn-back/infra/database/migrations"
)

func LoadConfigs() {
	setupTimezone()
	loadEnvVariables()
	database.Connect()
	migrations.Migrate()
}
