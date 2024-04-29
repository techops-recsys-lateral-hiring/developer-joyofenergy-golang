package priceplans

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

func callEndpoint(handler http.HandlerFunc, url string, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("request creation failed: %s", err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	return rr
}

func TestCompareAllPricePlansReturnResultFromService(t *testing.T) {
	s := &MockService{}
	h := NewHandler(s)
	params := httprouter.Params{{Key: "smartMeterId", Value: "123"}}
	compareAllPricePlans := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		h.CompareAll(writer, request, params)
	})

	rr := callEndpoint(compareAllPricePlans, "/price-plans/compare-all/123", t)
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned status code %v on valid request", rr.Code)
	var data domain.PricePlanComparisons

	err := json.Unmarshal(rr.Body.Bytes(), &data)
	assert.NoError(t, err)

	assert.Equal(t, domain.PricePlanComparisons{}, data)
}

func TestCompareAllPricePlansHandleServiceError(t *testing.T) {
	s := &MockService{err: errors.New("oops")}
	h := NewHandler(s)
	params := httprouter.Params{{Key: "smartMeterId", Value: "123"}}
	compareAllPricePlans := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		h.CompareAll(writer, request, params)
	})

	rr := callEndpoint(compareAllPricePlans, "/price-plans/compare-all/123", t)
	assert.NotEqual(t, http.StatusOK, rr.Code, "handler returned status code %v on failing request", rr.Code)

	var response domain.ErrorResponse

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, domain.ErrorResponse{Message: "oops"}, response)
}

func TestCompareAllPricePlansHandlerWithInvalidInput(t *testing.T) {
	s := &MockService{}
	h := NewHandler(s)
	params := httprouter.Params{{Key: "smartMeterId", Value: ""}}
	recommendPricePlans := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		h.Recommend(writer, request, params)
	})

	rr := callEndpoint(recommendPricePlans, "/price-plans/recommend/", t)
	assert.NotEqual(t, http.StatusOK, rr.Code, "handler returned status code %v on failing request", rr.Code)

	var response domain.ErrorResponse

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, domain.ErrorResponse{Message: "cannot be blank"}, response)
}

type MockService struct {
	err error
	Service
}

func (s *MockService) CompareAllPricePlans(smartMeterId string) (domain.PricePlanComparisons, error) {
	return domain.PricePlanComparisons{}, s.err
}
