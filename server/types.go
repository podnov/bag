package server

import "time"

type AccountTokenStatistics struct {
	AccountAddress string
	DaysSinceFirstTransaction float64
	Decimals int
	EarnedBalanceRatio float64
	EarnedTokenCount float64
	EarnedTokenCountPerDay float64
	EarnedTokenCountPerWeek float64
	EarnedValue float64
	EarnedValuePerDay float64
	EarnedValuePerWeek float64
	FirstTransactionTime time.Time
	TokenAddress string
	TokenCount float64
	TokenName string
	TokenPrice float64
	TokenPriceUpdatedAt time.Time
	Value float64
}
