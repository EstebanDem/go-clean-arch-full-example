package restclients

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPresetCurrencyConverterOk(t *testing.T) {
	converter := NewPresetCurrencyConverter()

	rate, err := converter.GetExchangeRate("USD", "ARS")

	assert.NoError(t, err)
	assert.Equal(t, 1300.00, rate)
}

func TestNewPresetCurrencyConverterBaseCurrencyNotFound(t *testing.T) {
	converter := NewPresetCurrencyConverter()

	rate, err := converter.GetExchangeRate("EUR", "ARS")

	assert.Equal(t, "exchange rate not found for EUR to ARS", err.Error())
	assert.Equal(t, 0.0, rate)
}

func TestNewPresetCurrencyConverterTargetCurrencyNotFound(t *testing.T) {
	converter := NewPresetCurrencyConverter()

	rate, err := converter.GetExchangeRate("USD", "EUR")

	assert.Equal(t, "exchange rate not found for USD to EUR", err.Error())
	assert.Equal(t, 0.0, rate)
}
