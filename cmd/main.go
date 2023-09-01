package main

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/config"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http"
)

func init() {
	config.LoadConfigs()
}

func main() {
	http.Server()
}
