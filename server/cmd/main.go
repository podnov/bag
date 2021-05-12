package main

import "encoding/json"
import "fmt"
import "os"

import "github.com/go-resty/resty/v2"
import "github.com/podnov/bag/server/bscscan"


const apiBaseUrl = "https://api.bscscan.com/api"
const apiKey = "X5T7BC9KUVXAWP6SQHRD5Z7RXRH58RJVIX" // TODO better

func main() {
	accountAddress := "0x5a6d55a598cba3a9fdafd0876c9ca02238c03e38"
	tokenAddress := "0x8076c74c5e3f5852037f31ff0093eeb8c8add8d3"

	balance, err := getAccountTokenBalance(accountAddress, tokenAddress)

	if err != nil {
		fmt.Println(err)
		fmt.Sprintf("Error: %v", fmt.Sprint(err))
		os.Exit(1)
	}

	fmt.Printf("Account balance: %d\n", balance)
	os.Exit(0)
}

func formatAccountTokenBalanceUrl(accountAddress string, tokenAddress string) string {
	return fmt.Sprintf("%s?module=account&action=tokenbalance&address=%s&contractaddress=%s&tag=latest&apikey=%s", apiBaseUrl, accountAddress, tokenAddress, apiKey)
}

func getAccountTokenBalance(accountAddress string, tokenAddress string) (int64, error) {
	url := formatAccountTokenBalanceUrl(accountAddress, tokenAddress)

	client := resty.New()

	resp, err := client.R().
		Get(url)

	if err != nil {
		return -1, err
	}

	apiResult := bscscan.ApiResult{}

	err = json.Unmarshal(resp.Body(), &apiResult)

	if err != nil {
		return -1, err
	}

	return apiResult.Result, nil
}
