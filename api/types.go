package api

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	PANCAKE_SWAP_V1 = "PANCAKE_SWAP_V1"
	PANCAKE_SWAP_V2 = "PANCAKE_SWAP_V2"
)

type AccountStatistics struct {
	AccountAddress string `json:"accountAddress"`
	AccruedValue decimal.Decimal `json:"accruedValue"`
	AccruedValuePerDay decimal.Decimal `json:"accruedValuePerDay"`
	AccruedValuePerWeek decimal.Decimal `json:"accruedValuePerWeek"`
	AccruedValueRatio decimal.Decimal `json:"accruedValueRatio"`
	FirstTransactionAt time.Time `json:"firstTransactionAt"`
	Tokens []AccountTokenStatistics `json:"tokens"`
	TransactionCount int `json:"transactionCount"`
	Value decimal.Decimal `json:"value"`
}

type AccountTokenStatistics struct {
	AccountAddress string `json:"accountAddress"`
	AccruedTokenCount decimal.Decimal `json:"accruedTokenCount"`
	AccruedTokenCountPerDay decimal.Decimal `json:"accruedTokenCountPerDay"`
	AccruedTokenCountPerWeek decimal.Decimal `json:"accruedTokenCountPerWeek"`
	AccruedValue decimal.Decimal `json:"accruedValue"`
	AccruedValuePerDay decimal.Decimal `json:"accruedValuePerDay"`
	AccruedValuePerWeek decimal.Decimal `json:"accruedValuePerWeek"`
	AccruedValueRatio decimal.Decimal `json:"accruedValueRatio"`
	CoinMarketCapId int `json:"coinMarketCapId"`
	DaysSinceFirstTransaction decimal.Decimal `json:"daysSinceFirstTransaction"`
	Decimals int `json:"decimals"`
	FirstTransactionAt time.Time `json:"firstTransactionAt"`
	TokenAddress string `json:"tokenAddress"`
	TokenCount decimal.Decimal `json:"tokenCount"`
	TokenName string `json:"tokenName"`
	TokenPrice decimal.Decimal `json:"tokenPrice"`
	TokenPriceSource string `json:"tokenPriceSource"`
	TokenPriceUpdatedAt time.Time `json:"tokenPriceUpdatedAt"`
	TransactionCount int `json:"transactionCount"`
	Value decimal.Decimal `json:"value"`
}

