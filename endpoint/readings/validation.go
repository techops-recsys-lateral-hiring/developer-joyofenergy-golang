package readings

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	validation "github.com/go-ozzo/ozzo-validation"

	"joi-energy-golang/domain"
)

func makeValidationMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			msg, ok := req.(domain.StoreReadings)
			if !ok {
				return nil, domain.ErrInvalidMessageType
			}
			if err := validateStoreReadings(msg); err != nil {
				return nil, fmt.Errorf("%w: %s", domain.ErrMissingArgument, err)
			}
			return next(ctx, req)
		}
	}
}

func validateStoreReadings(msg domain.StoreReadings) error {
	if err := validation.ValidateStruct(
		&msg,
		validation.Field(&msg.SmartMeterId, validation.Required),
		validation.Field(&msg.ElectricityReadings, validation.NotNil),
	); err != nil {
		return err
	}
	for _, row := range msg.ElectricityReadings {
		if err := validateElectricityReadings(row); err != nil {
			return err
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
