package readings

import (
	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

type Service interface {
	StoreReadings(smartMeterId string, reading []domain.ElectricityReading)
	GetReadings(smartMeterId string) []domain.ElectricityReading
}

type service struct {
	meterReadings *repository.MeterReadings
}

func NewService(
	meterReadings *repository.MeterReadings,
) Service {
	return &service{
		meterReadings: meterReadings,
	}
}

func (s *service) StoreReadings(smartMeterId string, reading []domain.ElectricityReading) {
	s.meterReadings.StoreReadings(smartMeterId, reading)
}

func (s *service) GetReadings(smartMeterId string) []domain.ElectricityReading {
	return s.meterReadings.GetReadings(smartMeterId)
}
