package services

type CurrencyConverter interface {
	GetExchangeRate(baseCurrency string, currency string) (float64, error)
}
