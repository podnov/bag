package server

import "math"
import "time"

import "github.com/podnov/bag/server/bscscan"
import "github.com/podnov/bag/server/pancakeswap"

type DataFetcher struct {
	bscClient *bscscan.BscApiClient
	pcsClient *pancakeswap.PancakeswapApiClient
}

func calculateEarnedTokens(balance int64, transactions *[]bscscan.TransactionApiResult) (int64) {
	for _, transaction := range *transactions {
		balance -= transaction.Value
	}

	return balance
}

func determineTokenDecimal(transactions *[]bscscan.TransactionApiResult) (int) {
	result := -1

	if len(*transactions) > 0 {
		result = (*transactions)[0].TokenDecimal
	}

	return result
}

func (df *DataFetcher) GetAccountTokenStatistics(accountAddress string, tokenAddress string) (*AccountTokenStatistics, error) {
	balance, err := df.bscClient.GetAccountTokenBalance(accountAddress, tokenAddress)

	if err != nil {
		return nil, err
	}

	transactions, err := df.bscClient.GetAccountTokenTransactions(accountAddress, tokenAddress)

	if err != nil {
		return nil, err
	}

	token, err := df.pcsClient.GetToken(tokenAddress)

	if err != nil {
		return nil, err
	}

	earned := calculateEarnedTokens(balance, &transactions)
	decimals := determineTokenDecimal(&transactions)

	divisor := math.Pow10(decimals)
	price := token.Data.Price
	priceUpdatedAt := time.Unix(0, token.UpdatedAt * int64(time.Millisecond))

	tokenCount := float64(balance) / divisor
	value := tokenCount * price
	earnedTokenCount := float64(earned) / divisor
	earnedValue := earnedTokenCount * price

	return &AccountTokenStatistics{
		Decimals: decimals,
		EarnedTokenCount: earnedTokenCount,
		EarnedValue: earnedValue,
		TokenCount: tokenCount,
		TokenPrice: price,
		TokenPriceUpdatedAt: priceUpdatedAt,
		Value: value,
	}, nil

}

func NewDataFetcher(bscClient *bscscan.BscApiClient, pcsClient *pancakeswap.PancakeswapApiClient) (*DataFetcher) {
	return &DataFetcher{
		bscClient: bscClient,
		pcsClient: pcsClient,
	}
}

