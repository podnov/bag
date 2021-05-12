package main

import "fmt"
import "os"

import "github.com/podnov/bag/server/bscscan"

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

	fmt.Printf("Account balance: %d\n", balance)
	os.Exit(0)
}

