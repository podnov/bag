package api

import "time"

const (
	PANCAKE_SWAP_V1 = "PANCAKE_SWAP_V1"
	PANCAKE_SWAP_V2 = "PANCAKE_SWAP_V2"
)

type AccountStatistics struct {
	AccountAddress string `json:"accountAddress"`
	EarnedValue float64 `json:"earnedValue"`
	EarnedValuePerDay float64 `json:"earnedValuePerDay"`
	EarnedValuePerWeek float64 `json:"earnedValuePerWeek"`
	EarnedValueRatio float64 `json:"earnedValueRatio"`
	FirstTransactionAt time.Time `json:"firstTransactionAt"`
	Tokens []AccountTokenStatistics `json:"tokens"`
	TransactionCount int `json:"transactionCount"`
	Value float64 `json:"value"`
}

type AccountTokenStatistics struct {
	AccountAddress string `json:"accountAddress"`
	DaysSinceFirstTransaction float64 `json:"daysSinceFirstTransaction"`
	Decimals int `json:"decimals"`
	EarnedTokenCount float64 `json:"earnedTokenCount"`
	EarnedTokenCountPerDay float64 `json:"earnedTokenCountPerDay"`
	EarnedTokenCountPerWeek float64 `json:"earnedTokenCountPerWeek"`
	EarnedValue float64 `json:"earnedValue"`
	EarnedValuePerDay float64 `json:"earnedValuePerDay"`
	EarnedValuePerWeek float64 `json:"earnedValuePerWeek"`
	EarnedValueRatio float64 `json:"earnedValueRatio"`
	FirstTransactionAt time.Time `json:"firstTransactionAt"`
	TokenAddress string `json:"tokenAddress"`
	TokenCount float64 `json:"tokenCount"`
	TokenName string `json:"tokenName"`
	TokenPrice float64 `json:"tokenPrice"`
	TokenPriceSource string `json:"tokenPriceSource"`
	TokenPriceUpdatedAt time.Time `json:"tokenPriceUpdatedAt"`
	TransactionCount int `json:"transactionCount"`
	Value float64 `json:"value"`
}

