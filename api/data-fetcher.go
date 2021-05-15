package api

import "errors"
import "fmt"
import "math"
import "math/big"
import "strings"
import "time"

import "github.com/podnov/bag/api/bscscan"
import pcsv1 "github.com/podnov/bag/api/pancakeswap/v1"
import pcsv2 "github.com/podnov/bag/api/pancakeswap/v2"

var daysPerWeek = big.NewFloat(float64(7))

type DataFetcher struct {
	bscClient *bscscan.BscApiClient
	pcsv1Client *pcsv1.PancakeswapApiClient
	pcsv2Client *pcsv2.PancakeswapApiClient
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
	price, priceUpdatedAtUnix, err := df.getTokenPrice(tokenAddress)

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
	priceUpdatedAt := time.Unix(0, priceUpdatedAtUnix * int64(time.Millisecond))
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
		EarnedTokenCount: earnedTokenCount,
		EarnedTokenCountPerDay: earnedTokenCountPerDay,
		EarnedTokenCountPerWeek: earnedTokenCountPerWeek,
		EarnedValue: earnedValue,
		EarnedValuePerDay: earnedValuePerDay,
		EarnedValuePerWeek: earnedValuePerWeek,
		EarnedValueRatio: earnedRatio,
		FirstTransactionAt: firstTransactionTime,
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
	var firstTransactionAt time.Time
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
			firstTransactionAt = tokenStatistics.FirstTransactionAt
		} else if tokenStatistics.FirstTransactionAt.Before(firstTransactionAt) {
			firstTransactionAt = tokenStatistics.FirstTransactionAt
		}

		tokens[tokenIndex] = tokenStatistics
		tokenIndex++
	}

	earnedValueRatio := new(big.Float).Quo(earnedValue, value)
	daysSinceFirstTransaction := big.NewFloat(time.Now().Sub(firstTransactionAt).Hours() / 24)
	earnedValuePerDay := new(big.Float).Quo(earnedValue, daysSinceFirstTransaction)
	earnedValuePerWeek := new(big.Float).Mul(earnedValuePerDay, daysPerWeek)

	return AccountStatistics{
		AccountAddress: accountAddress,
		EarnedValue: earnedValue,
		EarnedValuePerDay: earnedValuePerDay,
		EarnedValuePerWeek: earnedValuePerWeek,
		EarnedValueRatio: earnedValueRatio,
		FirstTransactionAt: firstTransactionAt,
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

func (df *DataFetcher) getTokenPrice(tokenAddress string) (*big.Float, int64, error) {
	v2Token, err := df.pcsv2Client.GetToken(tokenAddress)

	if err != nil {
		return nil, -1, err
	}

	price := v2Token.Data.Price
	priceUpdatedAt := v2Token.UpdatedAt

	if price == 0 {
		v1Token, err := df.pcsv1Client.GetToken(tokenAddress)

		if err != nil {
			return nil, -1, err
		}

		price = v1Token.Data.Price
		priceUpdatedAt = v1Token.UpdatedAt
	}

	return big.NewFloat(price), priceUpdatedAt, nil
}

func NewDataFetcher(bscClient *bscscan.BscApiClient, pcsv1Client *pcsv1.PancakeswapApiClient, pcsv2Client *pcsv2.PancakeswapApiClient) (DataFetcher) {
	return DataFetcher{
		bscClient: bscClient,
		pcsv1Client: pcsv1Client,
		pcsv2Client: pcsv2Client,
	}
}

func parseBigInt(value string) (*big.Int, error) {
	result, success := new(big.Int).SetString(value, 10)

	if !success {
		return nil, errors.New(fmt.Sprintf("Could not parse [%s] as big.Int", value))
	}

	return result, nil
}

