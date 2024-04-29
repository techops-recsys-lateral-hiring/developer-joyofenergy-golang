package readings

import (
	"testing"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

func TestStoreReadings(t *testing.T) {
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{},
	)
	service := NewService(
		&meterReadings,
	)
	service.StoreReadings("1", []domain.ElectricityReading{})
}
