package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

func TestEndpointEndpointSuccess(t *testing.T) {
	testHandler := setUpServer()

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/readings/read/smartMeterId", nil)
	req.Header.Add("Content-Type", "application/json")

	testHandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	result := rr.Result()

	expectedMessage := domain.StoreReadings{
		SmartMeterId: "smartMeterId",
		ElectricityReadings: nil,
	}
	expected, _ := json.MarshalIndent(expectedMessage, "", "  ")
	actual, err := ioutil.ReadAll(result.Body)
	_ = result.Body.Close()

	assert.NoError(t, err)
	assert.Equal(t, string(expected), string(actual))
}
