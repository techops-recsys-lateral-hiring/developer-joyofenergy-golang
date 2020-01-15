package priceplans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulValidation(t *testing.T) {
	err := validateSmartMeterId("smart-meter-0")
	assert.NoError(t, err)
}

func TestValidationFailureWithMissingID(t *testing.T) {
	err := validateSmartMeterId("")

	expectedErr := "cannot be blank"
	assert.EqualError(t, err, expectedErr)
}
