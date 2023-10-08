package config

import (
	"fmt"
	"time"
)

func setupTimezone() {
    time.Local, _ = time.LoadLocation("America/Sao_Paulo")
    currentTime := time.Now()
    fmt.Println("Current timezone:", time.Local)
    fmt.Println("Current time:", currentTime.Format(time.RFC3339))
}