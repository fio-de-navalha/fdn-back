package config

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database/migrations"
)

func LoadConfigs() {
	setupTimezone()
	loadEnvVariables()
	database.Connect()
	migrations.Migrate()
}
