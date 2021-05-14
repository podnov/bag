package server

import "math/big"
import "time"

type AccountStatistics struct {
	AccountAddress string
	EarnedValue *big.Float
	EarnedValuePerDay *big.Float
	EarnedValuePerWeek *big.Float
	EarnedValueRatio *big.Float
	FirstTransactionTime time.Time
	Tokens []AccountTokenStatistics
	TransactionCount int
	Value *big.Float
}

type AccountTokenStatistics struct {
	AccountAddress string
	DaysSinceFirstTransaction *big.Float
	Decimals *big.Int
	EarnedBalanceRatio *big.Float // TODO swich to earned value
	EarnedTokenCount *big.Float
	EarnedTokenCountPerDay *big.Float
	EarnedTokenCountPerWeek *big.Float
	EarnedValue *big.Float
	EarnedValuePerDay *big.Float
	EarnedValuePerWeek *big.Float
	FirstTransactionTime time.Time
	TokenAddress string
	TokenCount *big.Float
	TokenName string
	TokenPrice *big.Float
	TokenPriceUpdatedAt time.Time
	TransactionCount int
	Value *big.Float
}
