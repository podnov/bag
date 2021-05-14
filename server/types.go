package server

import "math/big"
import "time"

type AccountStatistics struct {
	AccountAddress string `json:"accountAddress"`
	EarnedValue *big.Float `json:"earnedValue"`
	EarnedValuePerDay *big.Float `json:"earnedValuePerDay"`
	EarnedValuePerWeek *big.Float `json:"earnedValuePerWeek"`
	EarnedValueRatio *big.Float `json:"earnedValueRatio"`
	FirstTransactionAt time.Time `json:"firstTransactionAt"`
	Tokens []AccountTokenStatistics `json:"tokens"`
	TransactionCount int `json:"transactionCount"`
	Value *big.Float `json:"value"`
}

type AccountTokenStatistics struct {
	AccountAddress string `json:"accountAddress"`
	DaysSinceFirstTransaction *big.Float `json:"daysSinceFirstTransaction"`
	Decimals *big.Int `json:"decimals"`
	EarnedTokenCount *big.Float `json:"earnedTokenCount"`
	EarnedTokenCountPerDay *big.Float `json:"earnedTokenCountPerDay"`
	EarnedTokenCountPerWeek *big.Float `json:"earnedTokenCountPerWeek"`
	EarnedValue *big.Float `json:"earnedValue"`
	EarnedValuePerDay *big.Float `json:"earnedValuePerDay"`
	EarnedValuePerWeek *big.Float `json:"earnedValuePerWeek"`
	EarnedValueRatio *big.Float `json:"earnedValueRatio"`
	FirstTransactionAt time.Time `json:"firstTransactionAt"`
	TokenAddress string `json:"tokenAddress"`
	TokenCount *big.Float `json:"tokenCount"`
	TokenName string `json:"tokenName"`
	TokenPrice *big.Float `json:"tokenPrice"`
	TokenPriceUpdatedAt time.Time `json:"tokenPriceUpdatedAt"`
	TransactionCount int `json:"transactionCount"`
	Value *big.Float `json:"value"`
}

