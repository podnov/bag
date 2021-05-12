package bscscan

type Int64ApiResult struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Result int64 `json:"result,string"`
}

type TransactionApiResult struct {
	BlockNumber int `json:"blockNumber,string"`
	TimeStamp int64 `json:"timeStamp,string"`
	Hash string `json:"hash"`
	Nonce int `json:"nonce,string"`
	BlockHash string `json:"blockHash"`
	From string `json:"from"`
	ContractAddress string `json:"contractAddress"`
	To string `json:"to"`
	Value int64 `json:"value,string"`
	TokenName string `json:"tokenName"`
	TokenSymbol string `json:"tokenSymbol"`
	TokenDecimal int `json:"tokenDecimal,string"`
	TransactionIndex int `json:"transactionIndex,string"`
	Gas int `json:"gas,string"`
	GasPrice int64 `json:"gasPrice,string"`
	GasUsed int `json:"gasUsed,string"`
	CumulativeGasUsed int `json:"cumulativeGasUsed,string"`
	Input string `json:"input"`
	Confirmations int `json:"confirmations,string"`
}

type TransactionsApiResult struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Result []TransactionApiResult `json:"result"`
}

