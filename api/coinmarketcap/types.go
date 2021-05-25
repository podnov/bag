package coinmarketcap

import (
	"time"
)

type StatusApiResult struct {
	Timestamp time.Time `json:"timestamp"`
	ErrorCode int `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed int `json:"elapsed"`
	CreditCount int `json:"credit_count"`
}

type CryptocurrencyMapEntryApiResult struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Slug string `json:"slug"`
	Rank int `json:"rank"`
	IsActive int `json:"is_active"`
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData time.Time `json:"last_historical_data"`
	Platform CryptocurrencyMapEntryPlatformApiResult `json:"platform"`
}

type CryptocurrencyMapEntryPlatformApiResult struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Slug string `json:"slug"`
	TokenAddress string `json:"token_address"`
}

type CryptocurrencyMapApiResult struct {
	Status StatusApiResult `json:"status"`
	Data []CryptocurrencyMapEntryApiResult `json:"data"`
}
