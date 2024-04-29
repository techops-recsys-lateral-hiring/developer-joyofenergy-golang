package priceplans

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

func validateSmartMeterId(smartMeterId string) error {
	return validation.Validate(smartMeterId, validation.Required)
}
