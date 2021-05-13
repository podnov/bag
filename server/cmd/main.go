package main

import "fmt"
import "os"

import "github.com/podnov/bag/server"
import "github.com/podnov/bag/server/bscscan"
import "github.com/podnov/bag/server/pancakeswap"

import "golang.org/x/text/message"

func main() {
	accountAddress := "0x5a6d55a598cba3a9fdafd0876c9ca02238c03e38"
	tokenAddress := "0x8076c74c5e3f5852037f31ff0093eeb8c8add8d3"

	bscClient := &bscscan.BscApiClient{}
	pcsClient := &pancakeswap.PancakeswapApiClient{}

	dataFetcher := server.NewDataFetcher(bscClient, pcsClient)

	statistics, err := dataFetcher.GetAccountTokenStatistics(accountAddress, tokenAddress)

	if err != nil {
		fmt.Println(err)
		fmt.Sprintf("Error: %v", fmt.Sprint(err))
		os.Exit(1)
	}

	printer := message.NewPrinter(message.MatchLanguage("en"))

	price := statistics.TokenPrice
	priceUpdatedAt := statistics.TokenPriceUpdatedAt
	tokenCount := printer.Sprintf("%f", statistics.TokenCount)
	value := printer.Sprintf("$%f", statistics.Value)
	earnedTokenCount := printer.Sprintf("%f", statistics.EarnedTokenCount)
	earnedValue := printer.Sprintf("$%f", statistics.EarnedValue)

	fmt.Printf("Token Price %.16f as of %s\n", price, priceUpdatedAt)
	fmt.Printf("Account balance: %s\n", tokenCount)
	fmt.Printf("Account balance value: %s\n", value)
	fmt.Printf("Earned tokens: %s\n", earnedTokenCount)
	fmt.Printf("Earned tokens value: %s\n", earnedValue)

	os.Exit(0)
}

