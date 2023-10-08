package config

import (
	"fmt"
	"time"
)

func setupTimezone() {
	location, err := time.LoadLocation("GMT-3")
	if err != nil {
		fmt.Println("Error loading timezone:", err)
		return
	}
    time.Local = location
    currentTime := time.Now()
    fmt.Println("Current time in GMT-3:", currentTime.Format(time.RFC3339))
}