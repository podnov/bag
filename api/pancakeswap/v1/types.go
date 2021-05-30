package v1

import (
	"github.com/shopspring/decimal"
)

type TokenApiResult struct {
	UpdatedAt int64 `json:"updated_at"`
	Data TokenDataApiResult `json:"data"`
}

type TokenDataApiResult struct {
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Price decimal.Decimal `json:"price,string"`
	PriceBNB decimal.Decimal `json:"price_BNB,string"`
}

