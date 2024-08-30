package services

import "github.com/stretchr/testify/mock"

type CurrencyConverterMock struct {
	mock.Mock
}

func (m *CurrencyConverterMock) GetExchangeRate(baseCurrency string, currency string) (float64, error) {
	args := m.Called(baseCurrency, currency)
	return args.Get(0).(float64), args.Error(1)
}
