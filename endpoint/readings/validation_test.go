package readings

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

func generateValidInput() domain.StoreReadings {
	return domain.StoreReadings{
		SmartMeterId: "12345",
		ElectricityReadings:         []domain.ElectricityReading{{
			Time: time.Now(),
			Reading: 123.45,
		}},
	}
}

func TestSuccessfulValidation(t *testing.T) {
	input := generateValidInput()

	err := validateStoreReadings(input)
	assert.NoError(t, err)
}

func TestValidationFailureWithMissingID(t *testing.T) {
	input := generateValidInput()
	input.SmartMeterId = ""

	err := validateStoreReadings(input)
	expectedErr := "smartMeterId: cannot be blank."
	assert.EqualError(t, err, expectedErr)
}

func TestValidationFailureWithMissingData(t *testing.T) {
	input := generateValidInput()
	input.ElectricityReadings = nil

	err := validateStoreReadings(input)
	expectedErr := "electricityReadings: is required."
	assert.EqualError(t, err, expectedErr)
}
