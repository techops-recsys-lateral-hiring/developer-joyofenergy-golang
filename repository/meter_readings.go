package repository

import "joi-energy-golang/domain"

type MeterReadings struct {
	meterAssociatedReadings map[string][]domain.ElectricityReading
}

func NewMeterReadings(meterAssociatedReadings map[string][]domain.ElectricityReading) MeterReadings {
	return MeterReadings{meterAssociatedReadings: meterAssociatedReadings}
}

func (m *MeterReadings) GetReadings(smartMeterId string) []domain.ElectricityReading {
	v, ok := m.meterAssociatedReadings[smartMeterId]
	if !ok {
		return nil
	}
	return v
}

func (m *MeterReadings) StoreReadings(smartMeterId string, electricityReadings []domain.ElectricityReading) {
	m.meterAssociatedReadings[smartMeterId] = append(m.meterAssociatedReadings[smartMeterId], electricityReadings...)
}
