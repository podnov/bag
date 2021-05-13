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

func determineFirstTransaction(transactions *[]bscscan.TransactionApiResult) (*bscscan.TransactionApiResult) {
	var result *bscscan.TransactionApiResult = nil

	if len(*transactions) > 0 {
		result = &(*transactions)[0]
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

	earned := calculateEarnedTokens(balance, transactions)
	firstTransaction := determineFirstTransaction(transactions)

	firstTransactionTime := time.Unix(firstTransaction.TimeStamp, 0)
	tokenName := firstTransaction.TokenName
	decimals := firstTransaction.TokenDecimal
	daysSinceFirstTransaction := time.Now().Sub(firstTransactionTime).Hours() / 24
	divisor := math.Pow10(decimals)
	price := token.Data.Price
	priceUpdatedAt := time.Unix(0, token.UpdatedAt * int64(time.Millisecond))

	tokenCount := float64(balance) / divisor
	value := tokenCount * price
	earnedTokenCount := float64(earned) / divisor
	earnedTokenCountPerDay := earnedTokenCount / daysSinceFirstTransaction
	earnedTokenCountPerWeek := earnedTokenCountPerDay * 7
	earnedValue := earnedTokenCount * price
	earnedValuePerDay := earnedTokenCountPerDay * price
	earnedValuePerWeek := earnedTokenCountPerWeek * price
	earnedRatio := earnedTokenCount / tokenCount

	return &AccountTokenStatistics{
		AccountAddress: accountAddress,
		DaysSinceFirstTransaction: daysSinceFirstTransaction,
		Decimals: decimals,
		EarnedBalanceRatio: earnedRatio,
		EarnedTokenCount: earnedTokenCount,
		EarnedTokenCountPerDay: earnedTokenCountPerDay,
		EarnedTokenCountPerWeek: earnedTokenCountPerWeek,
		EarnedValue: earnedValue,
		EarnedValuePerDay: earnedValuePerDay,
		EarnedValuePerWeek: earnedValuePerWeek,
		FirstTransactionTime: firstTransactionTime,
		TokenAddress: tokenAddress,
		TokenCount: tokenCount,
		TokenName: tokenName,
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

