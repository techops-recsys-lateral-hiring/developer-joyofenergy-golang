package readings

import (
	"bytes"
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
	Service
}

func (s *MockService) StoreReadings(smartMeterId string, reading []domain.ElectricityReading) {}

func TestMakeStoreReadingsHandler(t *testing.T) {
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeStoreReadingsHandler(mockService, mockLogger)
	r := httptest.NewRecorder()

	input := generateValidInput()
	buf := bytes.NewBuffer(nil)
	data, _ := json.MarshalIndent(&input, "", "  ")
	buf.Write(data)

	req := httptest.NewRequest("POST", "/Endpoint", buf)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusOK, actualStatusCode)
	err := result.Body.Close()
	assert.NoError(t, err)
}

func TestMakeStoreReadingsHandlerWithInvalidInput(t *testing.T) {
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeStoreReadingsHandler(mockService, mockLogger)
	r := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/Endpoint", nil)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusInternalServerError, actualStatusCode)

	expectedMessage := domain.Error{
		ErrorMessage: "unexpected end of JSON input",
	}
	expected, _ := json.MarshalIndent(expectedMessage, "", "  ")
	actual, err := ioutil.ReadAll(result.Body)
	_ = result.Body.Close()

	assert.NoError(t, err)
	assert.Equal(t, string(expected), string(actual))
}
