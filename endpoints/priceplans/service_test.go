package priceplans

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

func TestCompareAllPricePlans(t *testing.T) {
	accounts := repository.NewAccounts(map[string]string{"home-sweet-home": "test-plan"})
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{"home-sweet-home": {{
			Time:    time.Now(),
			Reading: 5.0,
		}, {
			Time:    time.Now().Add(-10 * time.Hour),
			Reading: 15.0,
		}}},
	)
	pricePlans := repository.NewPricePlans(
		[]domain.PricePlan{{
			PlanName: "test-plan",
			UnitRate: 3.0,
		}},
		&meterReadings,
	)
	service := NewService(
		&pricePlans,
		&accounts,
	)
	plans, err := service.CompareAllPricePlans("home-sweet-home")
	expected := domain.PricePlanComparisons{
		PricePlanId: "test-plan",
		PricePlanComparisons: map[string]float64{
			"test-plan": 3.0,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected.PricePlanId, plans.PricePlanId)
	assert.InDelta(t, expected.PricePlanComparisons["test-plan"], plans.PricePlanComparisons["test-plan"], 0.001)
}
