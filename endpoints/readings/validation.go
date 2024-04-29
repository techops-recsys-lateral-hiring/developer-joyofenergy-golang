package readings

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"

	"joi-energy-golang/domain"
)

func validateStoreReadings(msg domain.StoreReadings) error {
	if err := validation.ValidateStruct(
		&msg,
		validation.Field(&msg.SmartMeterId, validation.Required),
		validation.Field(&msg.ElectricityReadings, validation.NotNil),
	); err != nil {
		return fmt.Errorf("store readings validation failed: %w", err)
	}
	for _, row := range msg.ElectricityReadings {
		if err := validateElectricityReadings(row); err != nil {
			return fmt.Errorf("store readings validation failed for electricity reading: %w", err)
		}
	}
	return nil
}

func validateElectricityReadings(row domain.ElectricityReading) error {
	return nil
}

func validateSmartMeterId(smartMeterId string) error {
	return validation.Validate(smartMeterId, validation.Required)
}
