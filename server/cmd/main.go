package main

import "fmt"
import "math"
import "os"

import "github.com/podnov/bag/server/bscscan"
import "github.com/podnov/bag/server/pancakeswap"

import "golang.org/x/text/message"

func main() {
	accountAddress := "0x5a6d55a598cba3a9fdafd0876c9ca02238c03e38"
	tokenAddress := "0x8076c74c5e3f5852037f31ff0093eeb8c8add8d3"

	bscClient := bscscan.BscApiClient{}

	balance, err := bscClient.GetAccountTokenBalance(accountAddress, tokenAddress)

	if err != nil {
		fmt.Println(err)
		fmt.Sprintf("Error: %v", fmt.Sprint(err))
		os.Exit(1)
	}

	transactions, err := bscClient.GetAccountTokenTransactions(accountAddress, tokenAddress)

	if err != nil {
		fmt.Println(err)
		fmt.Sprintf("Error: %v", fmt.Sprint(err))
		os.Exit(1)
	}

	pcsClient := pancakeswap.PancakeswapApiClient{}

	token, err := pcsClient.GetToken(tokenAddress)

	if err != nil {
		fmt.Println(err)
		fmt.Sprintf("Error: %v", fmt.Sprint(err))
		os.Exit(1)
	}

	earned := calculateEarnedTokens(balance, &transactions)
	decimals := determineTokenDecimal(&transactions)
	price := token.Data.Price

	fmt.Printf("Price %.16f\n", price)

	divisor := math.Pow10(decimals)

	decimalBalance := float64(balance) / divisor
	decimalBalanceValue := decimalBalance * price
	decimalEarned := float64(earned) / divisor
	decimalEarnedValue := decimalEarned * price

	printer := message.NewPrinter(message.MatchLanguage("en"))

	formattedBalance := printer.Sprintf("%f", decimalBalance)
	formattedBalanceValue := printer.Sprintf("$%f", decimalBalanceValue)
	formattedEarned := printer.Sprintf("%f", decimalEarned)
	formattedEarnedValue := printer.Sprintf("$%f", decimalEarnedValue)

	fmt.Printf("Account balance: %s\n", formattedBalance)
	fmt.Printf("Account balance value: %s\n", formattedBalanceValue)
	fmt.Printf("Earned tokens: %s\n", formattedEarned)
	fmt.Printf("Earned tokens value: %s\n", formattedEarnedValue)

	os.Exit(0)
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
