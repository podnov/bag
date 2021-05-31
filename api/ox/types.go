package ox

import (
	"github.com/shopspring/decimal"
)

type QuoteApiResult struct {
	Price              decimal.Decimal        `json:"price"`
	GuaranteedPrice    decimal.Decimal        `json:"guaranteedPrice"`
	To                 string                 `json:"to"`
	Data               string                 `json:"data"`
	Value              string                 `json:"value"`
	Gas                int64                  `json:"gas,string"`
	EstimatedGas       int64                  `json:"estimatedGas,string"`
	GasPrice           decimal.Decimal        `json:"gasPrice"`
	ProtocolFee        int64                  `json:"protocolFee,string"`
	MinimumProtocolFee int64                  `json:"minimumProtocolFee,string"`
	BuyTokenAddress    string                 `json:"buyTokenAddress"`
	SellTokenAddress   string                 `json:"sellTokenAddress"`
	BuyAmount          string                 `json:"buyAmount"`
	SellAmount         string                 `json:"sellAmount"`
	AllowanceTarget    string                 `json:"allowanceTarget"`
	SellTokenToEthRate decimal.Decimal        `json:"sellTokenToEthRate"`
	BuyTokenToEthRate  decimal.Decimal        `json:"buyTokenToEthRate"`
	Sources            []QuoteSourceApiResult `json:"sources"`
	Orders             []QuoteOrderApiResult  `json:"orders"`
}

type QuoteOrderApiResult struct {
	MakerToken   string                      `json:"makerToken"`
	TakerToken   string                      `json:"takerToken"`
	MakerAmount  string                      `json:"makerAmount"`
	TakerAmount  string                      `json:"takerAmount"`
	FillData     QuoteOrderFillDataApiResult `json:"fillData"`
	Source       string                      `json:"source"`
	SourcePathId string                      `json:"sourcePathId"`
	Type         int                         `json:"type"`
}

type QuoteOrderFillDataApiResult struct {
	TokenAddressPath []string `json:"tokenAddressPath"`
	Router           string   `json:"router"`
}

type QuoteSourceApiResult struct {
	Name       string          `json:"name"`
	Proportion decimal.Decimal `json:"proportion"`
}
