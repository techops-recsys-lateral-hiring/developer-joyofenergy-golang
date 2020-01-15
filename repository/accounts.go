package repository

type Accounts struct {
	smartMeterToPricePlanAccounts map[string]string
}

func NewAccounts(smartMeterToPricePlanAccounts map[string]string) Accounts {
	return Accounts{
		smartMeterToPricePlanAccounts: smartMeterToPricePlanAccounts,
	}
}

func (a *Accounts) PricePlanIdForSmartMeterId(smartMeterId string) string {
	// TODO indicate missing value
	return a.smartMeterToPricePlanAccounts[smartMeterId]
}
