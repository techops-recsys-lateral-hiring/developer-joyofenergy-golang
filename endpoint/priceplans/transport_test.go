package priceplans

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

type MockService struct{
	err error
	Service
}

func (s *MockService) CompareAllPricePlans(smartMeterId string) (domain.PricePlanComparisons, error) {
	return domain.PricePlanComparisons{}, s.err
}

func TestMakeCompareAllPricePlanHandler(t *testing.T) {
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeCompareAllPricePlansHandler(mockService, mockLogger)
	r := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/price-plans/recommend/smart-meter-12345", nil)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusOK, actualStatusCode)
	err := result.Body.Close()
	assert.NoError(t, err)
}

func TestMakeCompareAllPricePlansHandlerWithInvalidInput(t *testing.T) {
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeCompareAllPricePlansHandler(mockService, mockLogger)
	r := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/price-plans/recommend/", nil)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusInternalServerError, actualStatusCode)

	expectedMessage := domain.Error{
		ErrorMessage: "cannot be blank",
	}
	expected, _ := json.MarshalIndent(expectedMessage, "", "  ")
	actual, err := ioutil.ReadAll(result.Body)
	_ = result.Body.Close()

	assert.NoError(t, err)
	assert.Equal(t, string(expected), string(actual))
}
