package api

import "time"

const (
	PANCAKE_SWAP_V1 = "PANCAKE_SWAP_V1"
	PANCAKE_SWAP_V2 = "PANCAKE_SWAP_V2"
)

type AccountStatistics struct {
	AccountAddress string `json:"accountAddress"`
	AccruedValue float64 `json:"accruedValue"`
	AccruedValuePerDay float64 `json:"accruedValuePerDay"`
	AccruedValuePerWeek float64 `json:"accruedValuePerWeek"`
	AccruedValueRatio float64 `json:"accruedValueRatio"`
	FirstTransactionAt time.Time `json:"firstTransactionAt"`
	Tokens []AccountTokenStatistics `json:"tokens"`
	TransactionCount int `json:"transactionCount"`
	Value float64 `json:"value"`
}

type AccountTokenStatistics struct {
	AccountAddress string `json:"accountAddress"`
	AccruedTokenCount float64 `json:"accruedTokenCount"`
	AccruedTokenCountPerDay float64 `json:"accruedTokenCountPerDay"`
	AccruedTokenCountPerWeek float64 `json:"accruedTokenCountPerWeek"`
	AccruedValue float64 `json:"accruedValue"`
	AccruedValuePerDay float64 `json:"accruedValuePerDay"`
	AccruedValuePerWeek float64 `json:"accruedValuePerWeek"`
	AccruedValueRatio float64 `json:"accruedValueRatio"`
	DaysSinceFirstTransaction float64 `json:"daysSinceFirstTransaction"`
	Decimals int `json:"decimals"`
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

