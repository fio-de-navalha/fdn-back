package config

import (
	migrations "github.com/fio-de-navalha/fdn-back/db"
	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
)

func LoadConfigs() {
	setupTimezone()
	loadEnvVariables()
	database.Connect()
	migrations.Migrate()
}
