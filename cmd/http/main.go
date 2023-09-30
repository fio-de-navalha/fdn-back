package main

import (
	"github.com/fio-de-navalha/fdn-back/internal/config"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http"
)

func init() {
	config.LoadConfigs()
	container.LoadContainers()
}

func main() {
	s := http.NewServer()
	s.StartServer()
}
