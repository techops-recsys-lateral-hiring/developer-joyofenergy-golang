package server

import (
	"math"
	"math/rand"
	"sort"
	"time"

	"joi-energy-golang/domain"
)

func defaultPricePlans() []domain.PricePlan {
	return []domain.PricePlan{
		{
			PlanName:       "price-plan-0",
			EnergySupplier: "Dr Evil's Dark Energy",
			UnitRate:       10,
		},
		{
			PlanName:       "price-plan-1",
			EnergySupplier: "The Green Eco",
			UnitRate:       2,
		},
		{
			PlanName:       "price-plan-2",
			EnergySupplier: "Power for Everyone",
			UnitRate:       1,
		},
	}
}

func defaultSmartMeterToPricePlanAccounts() map[string]string {
	return map[string]string {
		"smart-meter-0": "price-plan-0",
		"smart-meter-1": "price-plan-1",
		"smart-meter-2": "price-plan-0",
		"smart-meter-3": "price-plan-2",
		"smart-meter-4": "price-plan-1",
	}
}

func defaultMeterElectricityReadings() map[string][]domain.ElectricityReading {
	res := map[string][]domain.ElectricityReading{}
	for k := range defaultSmartMeterToPricePlanAccounts() {
		res[k] = generateElectricityReadings(20)
	}
	return res
}

func generateElectricityReadings(number int) []domain.ElectricityReading {
	readings := make([]domain.ElectricityReading, number)
	now := time.Now()
	for i := range readings {
		electricityReading := domain.ElectricityReading{
			Time:    now.Add(time.Duration(i*-10) * time.Second),
			Reading: math.Abs(rand.NormFloat64()),
		}
		readings[i] = electricityReading
	}
	sort.Slice(readings, func(i, j int) bool { return readings[i].Time.Before(readings[j].Time) })
	return readings
}
