package priceplans

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

func TestCompareAllPricePlansReturnResultFromService(t *testing.T) {
	s := &MockService{}
	e := makeCompareAllPricePlansEndpoint(s)

	response, err := e(context.Background(), "123")
	expectedResponse := domain.PricePlanComparisons{}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestCompareAllPricePlansHandleServiceError(t *testing.T) {
	s := &MockService{err: errors.New("oops")}
	e := makeCompareAllPricePlansEndpoint(s)

	_, err := e(context.Background(), "123")
	expectedErr := "oops"

	assert.EqualError(t, err, expectedErr)
}
