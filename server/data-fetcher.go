package server

import "errors"
import "fmt"
import "math"
import "math/big"
import "strings"
import "time"

import "github.com/podnov/bag/server/bscscan"
import "github.com/podnov/bag/server/pancakeswap"

var daysPerWeek = big.NewFloat(float64(7))

type DataFetcher struct {
	bscClient *bscscan.BscApiClient
	pcsClient *pancakeswap.PancakeswapApiClient
}

func calculateEarnedRawTokens(accountAddress string, balance *big.Int, transactions []bscscan.TransactionApiResult) (*big.Int, error) {
	result := new(big.Int).Set(balance)

	for _, transaction := range transactions {
		value, err := parseBigInt(transaction.Value)

		if err != nil {
			return nil, err
		}

		if strings.EqualFold(transaction.To, accountAddress) {
			result.Sub(result, value)
		} else {
			result.Add(result, value)
		}
	}

	return result, nil
}

func (df *DataFetcher) createAccountTokenStatistics(accountAddress string, tokenAddress string, transactions []bscscan.TransactionApiResult) (AccountTokenStatistics, error) {
	token, err := df.pcsClient.GetToken(tokenAddress)

	if err != nil {
		return AccountTokenStatistics{}, err
	}

	untypedRawBalance, err := df.bscClient.GetAccountTokenBalance(accountAddress, tokenAddress)

	if err != nil {
		return AccountTokenStatistics{}, err
	}

	rawBalance, err := parseBigInt(untypedRawBalance)

	if err != nil {
		return AccountTokenStatistics{}, err
	}

	rawEarned, err := calculateEarnedRawTokens(accountAddress, rawBalance, transactions)

	if err != nil {
		return AccountTokenStatistics{}, err
	}

	firstTransaction := determineFirstTransaction(transactions)

	firstTransactionTime := time.Unix(firstTransaction.TimeStamp, 0)
	tokenName := firstTransaction.TokenName
	decimals := firstTransaction.TokenDecimal
	daysSinceFirstTransaction := big.NewFloat(time.Now().Sub(firstTransactionTime).Hours() / 24)
	divisor := big.NewFloat(math.Pow10(decimals))
	price := big.NewFloat(token.Data.Price)
	priceUpdatedAt := time.Unix(0, token.UpdatedAt * int64(time.Millisecond))
	transactionCount := len(transactions)

	tokenCount := new(big.Float).Quo(new(big.Float).SetInt(rawBalance), divisor)
	value := new(big.Float).Mul(tokenCount, price)
	earnedTokenCount := new(big.Float).Quo(new(big.Float).SetInt(rawEarned), divisor)
	earnedTokenCountPerDay := new(big.Float).Quo(earnedTokenCount, daysSinceFirstTransaction)
	earnedTokenCountPerWeek := new(big.Float).Mul(earnedTokenCountPerDay, daysPerWeek)
	earnedValue := new(big.Float).Mul(earnedTokenCount, price)
	earnedValuePerDay := new(big.Float).Mul(earnedTokenCountPerDay, price)
	earnedValuePerWeek := new(big.Float).Mul(earnedTokenCountPerWeek, price)
	earnedRatio := new(big.Float).Quo(earnedTokenCount, tokenCount)

	return AccountTokenStatistics{
		AccountAddress: accountAddress,
		DaysSinceFirstTransaction: daysSinceFirstTransaction,
		Decimals: big.NewInt(int64(decimals)),
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
		TransactionCount: transactionCount,
		Value: value,
	}, nil
}

func determineFirstTransaction(transactions []bscscan.TransactionApiResult) (bscscan.TransactionApiResult) {
	var result bscscan.TransactionApiResult

	if len(transactions) > 0 {
		result = transactions[0]
	} else {
		result = bscscan.TransactionApiResult{}
	}

	return result
}

func mapTokenTransactions(transactions []bscscan.TransactionApiResult) (map[string][]bscscan.TransactionApiResult) {
	result := make(map[string][]bscscan.TransactionApiResult)

	for _, transaction := range transactions {
		tokenAddress := transaction.ContractAddress
		tokenTransactions, exists := result[tokenAddress];

		if exists {
			tokenTransactions = append(tokenTransactions, transaction)
		} else {
			tokenTransactions = []bscscan.TransactionApiResult{transaction}
		}

		result[tokenAddress] = tokenTransactions
	}

	return result
}

func (df *DataFetcher) GetAccountStatistics(accountAddress string) (AccountStatistics, error) {
	transactions, err := df.bscClient.GetAccountTokenTransactions(accountAddress)

	if err != nil {
		return AccountStatistics{}, err
	}

	transactionsByToken := mapTokenTransactions(transactions)
	tokens := make([]AccountTokenStatistics, len(transactionsByToken))

	tokenIndex := 0
	earnedValue := big.NewFloat(float64(0))
	var firstTransactionTime time.Time
	transactionCount := 0
	value := big.NewFloat(float64(0))

	for tokenAddress, tokenTransactions := range transactionsByToken {
		tokenStatistics, err := df.createAccountTokenStatistics(accountAddress, tokenAddress, tokenTransactions)

		if err != nil {
			return AccountStatistics{}, err
		}

		earnedValue.Add(earnedValue, tokenStatistics.EarnedValue)
		transactionCount += tokenStatistics.TransactionCount
		value.Add(value, tokenStatistics.Value)

		if tokenIndex == 0 {
			firstTransactionTime = tokenStatistics.FirstTransactionTime
		} else if tokenStatistics.FirstTransactionTime.Before(firstTransactionTime) {
			firstTransactionTime = tokenStatistics.FirstTransactionTime
		}

		tokens[tokenIndex] = tokenStatistics
		tokenIndex++
	}

	earnedValueRatio := new(big.Float).Quo(earnedValue, value)
	daysSinceFirstTransaction := big.NewFloat(time.Now().Sub(firstTransactionTime).Hours() / 24)
	earnedValuePerDay := new(big.Float).Quo(earnedValue, daysSinceFirstTransaction)
	earnedValuePerWeek := new(big.Float).Mul(earnedValuePerDay, daysPerWeek)

	return AccountStatistics{
		AccountAddress: accountAddress,
		EarnedValue: earnedValue,
		EarnedValuePerDay: earnedValuePerDay,
		EarnedValuePerWeek: earnedValuePerWeek,
		EarnedValueRatio: earnedValueRatio,
		FirstTransactionTime: firstTransactionTime,
		Tokens: tokens,
		TransactionCount: transactionCount,
		Value: value,
	}, nil
}

func (df *DataFetcher) GetAccountStatisticsForToken(accountAddress string, tokenAddress string) (AccountTokenStatistics, error) {
	transactions, err := df.bscClient.GetAccountTokenTransactionsForToken(accountAddress, tokenAddress)

	if err != nil {
		return AccountTokenStatistics{}, err
	}

	return df.createAccountTokenStatistics(accountAddress, tokenAddress, transactions)
}

func NewDataFetcher(bscClient *bscscan.BscApiClient, pcsClient *pancakeswap.PancakeswapApiClient) (DataFetcher) {
	return DataFetcher{
		bscClient: bscClient,
		pcsClient: pcsClient,
	}
}

func parseBigInt(value string) (*big.Int, error) {
	result, success := new(big.Int).SetString(value, 10)

	if !success {
		return nil, errors.New(fmt.Sprintf("Could not parse [%s] as big.Int", value))
	}

	return result, nil
}

