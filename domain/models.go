package domain

import (
	"time"
)

type ElectricityReading struct {
	Time    time.Time `json:"time"`
	Reading float64   `json:"reading"`
}

type PricePlan struct {
	EnergySupplier      string
	PlanName            string
	UnitRate            float64
	PeakTimeMultipliers []PeakTimeMultiplier
}

type PeakTimeMultiplier struct {
	DayOfWeek  time.Weekday
	Multiplier float64
}

type SingleRecommendation struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type PricePlanRecommendation struct {
	Recommendations []SingleRecommendation `json:"recommendations"`
}

type PricePlanComparisons struct {
	PricePlanId          string             `json:"pricePlanId"`
	PricePlanComparisons map[string]float64 `json:"pricePlanComparisons"`
}

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type StoreReadings struct {
	SmartMeterId        string               `json:"smartMeterId"`
	ElectricityReadings []ElectricityReading `json:"electricityReadings"`
}
