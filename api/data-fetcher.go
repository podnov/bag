package api

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/podnov/bag/api/bscscan"
	"github.com/podnov/bag/api/coinmarketcap"
	pcsv1 "github.com/podnov/bag/api/pancakeswap/v1"
	pcsv2 "github.com/podnov/bag/api/pancakeswap/v2"
)

var daysPerWeek = decimal.NewFromInt(7)
var hoursPerDay = decimal.NewFromInt(24)
var millisPerHour = decimal.NewFromInt(3600000)
var ten = decimal.NewFromInt(10)

type DataFetcher struct {
	bscClient              *bscscan.BscApiClient
	cryptocurrencyMapStore *coinmarketcap.CryptocurrencyMapStore
	pcsv1Client            *pcsv1.PancakeswapApiClient
	pcsv2Client            *pcsv2.PancakeswapApiClient
}

func calculateAccruedRawTokens(accountAddress string, balance *big.Int, transactions []bscscan.TransactionApiResult) (*big.Int, error) {
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

func calculateDaysPriorToNow(firstTransactionTime time.Time) decimal.Decimal {
	millis := time.Now().Sub(firstTransactionTime).Milliseconds()
	return decimal.NewFromInt(millis).
		Div(millisPerHour).
		Div(hoursPerDay)
}

func (df *DataFetcher) createAccountTokenStatistics(accountAddress string, tokenAddress string, transactions []bscscan.TransactionApiResult) (AccountTokenStatistics, error) {
	price, priceUpdatedAtUnix, priceSource, err := df.getTokenPrice(tokenAddress)

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

	rawAccrued, err := calculateAccruedRawTokens(accountAddress, rawBalance, transactions)

	if err != nil {
		return AccountTokenStatistics{}, err
	}

	firstTransaction := determineFirstTransaction(transactions)

	coinMarketCapId := df.cryptocurrencyMapStore.GetEntry(firstTransaction.TokenSymbol).Id
	firstTransactionTime := time.Unix(firstTransaction.TimeStamp, 0)
	tokenName := firstTransaction.TokenName
	priceUpdatedAt := time.Unix(0, priceUpdatedAtUnix*int64(time.Millisecond))
	transactionCount := len(transactions)

	decimals := firstTransaction.TokenDecimal
	divisor := ten.Pow(decimal.NewFromInt(int64(decimals)))
	decimalBalance := decimal.NewFromBigInt(rawBalance, 0)
	decimalAccrued := decimal.NewFromBigInt(rawAccrued, 0)

	tokenCount := decimalBalance.Div(divisor)
	accruedTokenCount := decimalAccrued.Div(divisor)

	value := tokenCount.Mul(price)
	accruedValue := accruedTokenCount.Mul(price)
	daysSinceFirstTransaction := calculateDaysPriorToNow(firstTransactionTime)

	accruedTokenCountPerDay := accruedTokenCount.Div(daysSinceFirstTransaction)
	accruedValuePerDay := accruedTokenCountPerDay.Mul(price)
	accruedTokenCountPerWeek := accruedTokenCountPerDay.Mul(daysPerWeek)
	accruedValuePerWeek := accruedValuePerDay.Mul(daysPerWeek)

	accruedRatio := accruedTokenCount.Div(tokenCount)

	return AccountTokenStatistics{
		AccountAddress:            accountAddress,
		AccruedTokenCount:         accruedTokenCount,
		AccruedTokenCountPerDay:   accruedTokenCountPerDay,
		AccruedTokenCountPerWeek:  accruedTokenCountPerWeek,
		AccruedValue:              accruedValue,
		AccruedValuePerDay:        accruedValuePerDay,
		AccruedValuePerWeek:       accruedValuePerWeek,
		AccruedValueRatio:         accruedRatio,
		CoinMarketCapId:           coinMarketCapId,
		DaysSinceFirstTransaction: daysSinceFirstTransaction,
		Decimals:                  decimals,
		FirstTransactionAt:        firstTransactionTime,
		TokenAddress:              tokenAddress,
		TokenCount:                tokenCount,
		TokenName:                 tokenName,
		TokenPrice:                price,
		TokenPriceSource:          priceSource,
		TokenPriceUpdatedAt:       priceUpdatedAt,
		TransactionCount:          transactionCount,
		Value:                     value,
	}, nil
}

func determineFirstTransaction(transactions []bscscan.TransactionApiResult) bscscan.TransactionApiResult {
	var result bscscan.TransactionApiResult

	if len(transactions) > 0 {
		result = transactions[0]
	} else {
		result = bscscan.TransactionApiResult{}
	}

	return result
}

func mapTokenTransactions(transactions []bscscan.TransactionApiResult) map[string][]bscscan.TransactionApiResult {
	result := make(map[string][]bscscan.TransactionApiResult)

	for _, transaction := range transactions {
		tokenAddress := transaction.ContractAddress
		tokenTransactions, exists := result[tokenAddress]

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
	accruedValue := decimal.Zero
	var firstTransactionAt time.Time
	transactionCount := 0
	value := decimal.Zero

	for tokenAddress, tokenTransactions := range transactionsByToken {
		tokenStatistics, err := df.createAccountTokenStatistics(accountAddress, tokenAddress, tokenTransactions)

		if err != nil {
			return AccountStatistics{}, err
		}

		accruedValue = accruedValue.Add(tokenStatistics.AccruedValue)
		transactionCount += tokenStatistics.TransactionCount
		value = value.Add(tokenStatistics.Value)

		if tokenIndex == 0 {
			firstTransactionAt = tokenStatistics.FirstTransactionAt
		} else if tokenStatistics.FirstTransactionAt.Before(firstTransactionAt) {
			firstTransactionAt = tokenStatistics.FirstTransactionAt
		}

		tokens[tokenIndex] = tokenStatistics
		tokenIndex++
	}

	accruedValueRatio := accruedValue.Div(value)

	daysSinceFirstTransaction := calculateDaysPriorToNow(firstTransactionAt)
	accruedValuePerDay := accruedValue.Div(daysSinceFirstTransaction)
	accruedValuePerWeek := accruedValuePerDay.Mul(daysPerWeek)

	return AccountStatistics{
		AccountAddress:      accountAddress,
		AccruedValue:        accruedValue,
		AccruedValuePerDay:  accruedValuePerDay,
		AccruedValuePerWeek: accruedValuePerWeek,
		AccruedValueRatio:   accruedValueRatio,
		FirstTransactionAt:  firstTransactionAt,
		Tokens:              tokens,
		TransactionCount:    transactionCount,
		Value:               value,
	}, nil
}

func (df *DataFetcher) GetAccountStatisticsForToken(accountAddress string, tokenAddress string) (AccountTokenStatistics, error) {
	transactions, err := df.bscClient.GetAccountTokenTransactionsForToken(accountAddress, tokenAddress)

	if err != nil {
		return AccountTokenStatistics{}, err
	}

	return df.createAccountTokenStatistics(accountAddress, tokenAddress, transactions)
}

func (df *DataFetcher) getTokenPrice(tokenAddress string) (decimal.Decimal, int64, string, error) {
	v2Token, err := df.pcsv2Client.GetToken(tokenAddress)
	source := ""

	if err != nil {
		return decimal.Zero, -1, source, err
	}

	price := v2Token.Data.Price
	priceUpdatedAt := v2Token.UpdatedAt

	if price.IsZero() {
		v1Token, err := df.pcsv1Client.GetToken(tokenAddress)

		if err != nil {
			return decimal.Zero, -1, source, err
		}

		price = v1Token.Data.Price
		priceUpdatedAt = v1Token.UpdatedAt
		source = PANCAKE_SWAP_V1
	} else {
		source = PANCAKE_SWAP_V2
	}

	return price, priceUpdatedAt, source, nil
}

func NewDataFetcher(bscClient *bscscan.BscApiClient,
	cryptocurrencyMapStore *coinmarketcap.CryptocurrencyMapStore,
	pcsv1Client *pcsv1.PancakeswapApiClient,
	pcsv2Client *pcsv2.PancakeswapApiClient) DataFetcher {

	return DataFetcher{
		bscClient:              bscClient,
		cryptocurrencyMapStore: cryptocurrencyMapStore,
		pcsv1Client:            pcsv1Client,
		pcsv2Client:            pcsv2Client,
	}
}

func parseBigInt(value string) (*big.Int, error) {
	result, success := new(big.Int).SetString(value, 10)

	if !success {
		return nil, errors.New(fmt.Sprintf("Could not parse [%s] as big.Int", value))
	}

	return result, nil
}
