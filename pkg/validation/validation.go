package validation

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

func ValidateDatetime(datetimeStr string, prefix string) error {
	layout := "15:04"
	parsedTime, err := time.Parse(layout, datetimeStr)
	if err != nil {
		return err
	}
	if parsedTime.Hour() < 1 || parsedTime.Hour() >= 23 {
		return errors.New(prefix + " time format is invalid")
	}
	if parsedTime.Minute() < 0 || parsedTime.Minute() >= 59 {
		return errors.New(prefix + " time format is invalid")
	}
	return nil
}

func ValidUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}
	return nil
}
