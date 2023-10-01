package validation

import (
	"errors"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/google/uuid"
)

func ValidateDatetime(datetimeStr string, prefix string) error {
	parsedTime, err := time.Parse(constants.OpenCloseLayout, datetimeStr)
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
