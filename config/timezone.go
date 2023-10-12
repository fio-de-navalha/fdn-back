package config

import (
	"fmt"
	"time"
)

var GMTMinus3 *time.Location

func setupTimezone() {
    time.Local, _ = time.LoadLocation("America/Sao_Paulo")
    GMTMinus3 = time.Local
    currentTime := time.Now()
    fmt.Println("Current timezone:", time.Local)
    fmt.Println("Current time:", currentTime.Format(time.RFC3339))
}