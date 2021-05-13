package server

import "time"

type AccountTokenStatistics struct {
	Decimals int
	EarnedTokenCount float64
	EarnedValue float64
	TokenCount float64
	TokenPrice float64
	TokenPriceUpdatedAt time.Time
	Value float64
}
