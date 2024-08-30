package restclients

import "fmt"

// PresetCurrencyConverter represents an external client that gives exchanges rates
type PresetCurrencyConverter struct {
	rates map[string]map[string]float64
}

func NewPresetCurrencyConverter() PresetCurrencyConverter {
	return PresetCurrencyConverter{rates: map[string]map[string]float64{
		"USD": {
			"MXN": 20.0,
			"ARS": 1300.0,
		},
	}}
}

func (p PresetCurrencyConverter) GetExchangeRate(baseCurrency string, currency string) (float64, error) {
	if targetRates, ok := p.rates[baseCurrency]; ok {
		if rate, ok := targetRates[currency]; ok {
			return rate, nil
		}
	}
	return 0, fmt.Errorf("exchange rate not found for %s to %s", baseCurrency, currency)
}
