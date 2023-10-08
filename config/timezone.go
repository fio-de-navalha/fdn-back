package config

import (
	"fmt"
	"time"
)

func setupTimezone() {
	location := time.FixedZone("GMT-3", -3*60*60)
    time.Local = location
    currentTime := time.Now()
    fmt.Println("Current time in GMT-3:", currentTime.Format(time.RFC3339))
}