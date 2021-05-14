package main

import "fmt"
import "math/big"
import "os"

import "github.com/podnov/bag/server"
import "github.com/podnov/bag/server/bscscan"
import "github.com/podnov/bag/server/pancakeswap"

import "golang.org/x/text/message"

var oneHundred = big.NewFloat(float64(100))

func main() {
	accountAddress := "0x5a6d55a598cba3a9fdafd0876c9ca02238c03e38"
	//accountAddress = "0x378Ec8Be66FD1EeAC595009c37A83e5c446EE146"

	bscClient := &bscscan.BscApiClient{}
	pcsClient := &pancakeswap.PancakeswapApiClient{}

	dataFetcher := server.NewDataFetcher(bscClient, pcsClient)

	statistics, err := dataFetcher.GetAccountStatistics(accountAddress)

	if err != nil {
		fmt.Sprintf("Error: %v", fmt.Sprint(err))
		os.Exit(1)
	}

	printer := message.NewPrinter(message.MatchLanguage("en"))

	tokens := statistics.Tokens
	tokenCount := len(tokens)

	for _, tokenStatistics := range tokens {
		printTokenStatistics(printer, tokenStatistics)
		fmt.Println("")
	}

	fmt.Printf("=== Summary ===\n")
	fmt.Printf("AccountAddress: %v\n", statistics.AccountAddress)
	fmt.Printf("Tokens held: %v\n", tokenCount)
	fmt.Printf("Transaction count: %v\n", statistics.TransactionCount)
	fmt.Printf("First transaction date: %s\n", statistics.FirstTransactionTime)
	fmt.Printf("Account value: %v\n", printer.Sprintf("$%f", statistics.Value))
	fmt.Printf("Earned value: %v\n", printer.Sprintf("$%f", statistics.EarnedValue))
	fmt.Printf("Earned percent of value: %.2f%%\n", new(big.Float).Mul(statistics.EarnedValueRatio, oneHundred))
	fmt.Printf("Earned value per day: %v\n", printer.Sprintf("$%f", statistics.EarnedValuePerDay))
	fmt.Printf("Earned value per week: %v\n", printer.Sprintf("$%f", statistics.EarnedValuePerWeek))

	os.Exit(0)
}

func printTokenStatistics(printer *message.Printer, statistics server.AccountTokenStatistics) {
	tokenAddress := statistics.TokenAddress
	tokenName := statistics.TokenName
	price := statistics.TokenPrice
	priceUpdatedAt := statistics.TokenPriceUpdatedAt
	tokenCount := printer.Sprintf("%f", statistics.TokenCount)
	value := printer.Sprintf("$%f", statistics.Value)
	earnedTokenCount := printer.Sprintf("%f", statistics.EarnedTokenCount)
	earnedTokenCountPerDay := printer.Sprintf("%f", statistics.EarnedTokenCountPerDay)
	earnedTokenCountPerWeek := printer.Sprintf("%f", statistics.EarnedTokenCountPerWeek)
	earnedValue := printer.Sprintf("$%f", statistics.EarnedValue)
	earnedValuePerDay := printer.Sprintf("$%f", statistics.EarnedValuePerDay)
	earnedValuePerWeek := printer.Sprintf("$%f", statistics.EarnedValuePerWeek)
	earnedBalanceRatio := printer.Sprintf("%.2f%%", new(big.Float).Mul(statistics.EarnedBalanceRatio, oneHundred))
	firstTransactionDate := statistics.FirstTransactionTime
	transactionCount := statistics.TransactionCount

	fmt.Printf("Account information for token %s (%s)\n", tokenName, tokenAddress)
	fmt.Printf("Token Price %.16f as of %s\n", price, priceUpdatedAt)
	fmt.Printf("Transaction count: %v\n", transactionCount)
	fmt.Printf("First transaction date: %s\n", firstTransactionDate)
	fmt.Printf("Account tokens balance: %s\n", tokenCount)
	fmt.Printf("Account tokens value: %s\n", value)
	fmt.Printf("Earned tokens: %s\n", earnedTokenCount)
	fmt.Printf("Earned tokens value: %s\n", earnedValue)
	fmt.Printf("Earned tokens percent of balance: %s\n", earnedBalanceRatio)
	fmt.Printf("Earned tokens per day: %s\n", earnedTokenCountPerDay)
	fmt.Printf("Earned tokens value per day: %s\n", earnedValuePerDay)
	fmt.Printf("Earned tokens per week: %s\n", earnedTokenCountPerWeek)
	fmt.Printf("Earned tokens value per week: %s\n", earnedValuePerWeek)
}

