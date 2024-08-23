package restclients

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

var (
	ErrCallingClient = errors.New("error calling free currency api client")
)

type FreeCurrencyApiClient struct {
	restClient *resty.Client
	apikey     string
}

func NewFreeCurrencyApiClient(apiKey string) *FreeCurrencyApiClient {
	return &FreeCurrencyApiClient{
		restClient: resty.New(),
		apikey:     apiKey,
	}
}

type FreeCurrencyApiJsonResponse struct {
	Data map[string]float64 `json:"data"`
}

func (f *FreeCurrencyApiClient) GetExchangeRate(baseCurrency string, currency string) (float64, error) {
	url := "https://api.freecurrencyapi.com/v1/latest?apikey=%s&currencies=%s&base_currency=%s"

	resp, err := f.restClient.R().Get(fmt.Sprintf(url, f.apikey, strings.ToUpper(currency), strings.ToUpper(baseCurrency)))

	if err != nil {
		return 0, err
	}

	if resp.IsError() {
		return 0, ErrCallingClient
	}

	var apiResponse FreeCurrencyApiJsonResponse
	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		return 0, err
	}

	return apiResponse.Data[currency], nil

}
