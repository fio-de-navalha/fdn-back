package config

import (
	"github.com/spf13/viper"
)

var (
	Port                string
	CloudFlareAccountId string
	CloudFlareReadToken string
	CloudFlareEditToken string
	CloudFlareBaseURL   string
)

func loadEnvVariables() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load environment variables")
	}

	Port = viper.GetString("PORT")
	CloudFlareAccountId = viper.GetString("CLOUDFLARE_ACCOUNT_ID")
	CloudFlareReadToken = viper.GetString("CLOUDFLARE_IMAGES_TOKEN")
	CloudFlareEditToken = viper.GetString("CLOUDFLARE_IMAGES_EDIT_TOKEN")
	CloudFlareBaseURL = "https://api.cloudflare.com/client/v4/accounts/" + CloudFlareAccountId + "/images/v1"
}
