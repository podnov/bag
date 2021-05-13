package main

import "fmt"
import "math"
import "os"

import "github.com/podnov/bag/server/bscscan"
import "golang.org/x/text/message"

func main() {
	accountAddress := "0x5a6d55a598cba3a9fdafd0876c9ca02238c03e38"
	tokenAddress := "0x8076c74c5e3f5852037f31ff0093eeb8c8add8d3"

	client := bscscan.BscApiClient{}

	balance, err := client.GetAccountTokenBalance(accountAddress, tokenAddress)

	if err != nil {
		fmt.Println(err)
		fmt.Sprintf("Error: %v", fmt.Sprint(err))
		os.Exit(1)
	}

	transactions, err := client.GetAccountTokenTransactions(accountAddress, tokenAddress)

	if err != nil {
		fmt.Println(err)
		fmt.Sprintf("Error: %v", fmt.Sprint(err))
		os.Exit(1)
	}

	earned := calculateEarnedTokens(balance, &transactions)
	decimals := determineTokenDecimal(&transactions)
	divisor := math.Pow10(decimals)

	printer := message.NewPrinter(message.MatchLanguage("en"))

	formattedBalance := printer.Sprintf("%f", float64(balance) / divisor)
	formattedEarned := printer.Sprintf("%f", float64(earned) / divisor)

	fmt.Printf("Account balance: %s\n", formattedBalance)
	fmt.Printf("Earned tokens: %s\n", formattedEarned)

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
