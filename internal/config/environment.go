package config

import (
	"github.com/spf13/viper"
)

func loadEnvVariables() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load environment variables")
	}
}
