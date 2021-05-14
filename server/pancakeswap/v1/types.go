package v1

type TokenApiResult struct {
	UpdatedAt int64 `json:"updated_at"`
	Data TokenDataApiResult `json:"data"`
}

type TokenDataApiResult struct {
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Price float64 `json:"price,string"`
	PriceBNB float64 `json:"price_BNB,string"`
}

