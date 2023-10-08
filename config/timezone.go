package config

import (
	"log"
	"time"
)

func setupTimezone() {
    local, err := time.LoadLocation("America/Sao_Paulo")
    if err != nil {
        log.Println(err)
    }
    
    time.Local = local
    currentTime := time.Now()
    log.Println("Current timezone:", time.Local)
    log.Println("Current time:", currentTime.Format(time.RFC3339))
}