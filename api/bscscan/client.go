package bscscan

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/podnov/bag/api/utils"
)

const ApiKeyEnvironmentVariableName = "BSC_SCAN_API_KEY"

const apiBaseUrl = "https://api.bscscan.com/api"

// TODO handle client invocation failures

type BscApiClient struct {

}

func (c *BscApiClient) createRestyClient() (*resty.Client) {
	return resty.New()
}

func (c *BscApiClient) formatAccountTokenBalanceUrl(accountAddress string, tokenAddress string) (string) {
	apiKey := utils.GetEnv(ApiKeyEnvironmentVariableName)
	return fmt.Sprintf("%s?module=account&action=tokenbalance&address=%s&contractaddress=%s&tag=latest&apikey=%s", apiBaseUrl, accountAddress, tokenAddress, apiKey)
}

func (c *BscApiClient) formatAccountTokenTransactionsUrl(accountAddress string) (string) {
	apiKey := utils.GetEnv(ApiKeyEnvironmentVariableName)
	return fmt.Sprintf("%s?module=account&action=tokentx&address=%s&sort=asc&apikey=%s", apiBaseUrl, accountAddress, apiKey)
}

func (c *BscApiClient) formatAccountTokenTransactionsForTokenUrl(accountAddress string, tokenAddress string) (string) {
	apiKey := utils.GetEnv(ApiKeyEnvironmentVariableName)
	return fmt.Sprintf("%s?module=account&action=tokentx&address=%s&contractaddress=%s&sort=asc&apikey=%s", apiBaseUrl, accountAddress, tokenAddress, apiKey)
}

func (c *BscApiClient) GetAccountTokenBalance(accountAddress string, tokenAddress string) (string, error) {
	url := c.formatAccountTokenBalanceUrl(accountAddress, tokenAddress)

	client := c.createRestyClient()

	resp, err := client.R().
		Get(url)

	if err != nil {
		return "", err
	}

	apiResult := StringApiResult{}

	err = json.Unmarshal(resp.Body(), &apiResult)

	if err != nil {
		return "", err
	}

	return apiResult.Result, nil
}

func (c *BscApiClient) GetAccountTokenTransactions(accountAddress string) ([]TransactionApiResult, error) {
	url := c.formatAccountTokenTransactionsUrl(accountAddress)

	client := c.createRestyClient()

	resp, err := client.R().
		Get(url)

	if err != nil {
		return nil, err
	}

	apiResult := TransactionsApiResult{}

	err = json.Unmarshal(resp.Body(), &apiResult)

	if err != nil {
		return nil, err
	}

	return apiResult.Result, nil
}

func (c *BscApiClient) GetAccountTokenTransactionsForToken(accountAddress string, tokenAddress string) ([]TransactionApiResult, error) {
	url := c.formatAccountTokenTransactionsForTokenUrl(accountAddress, tokenAddress)

	client := c.createRestyClient()

	resp, err := client.R().
		Get(url)

	if err != nil {
		return nil, err
	}

	apiResult := TransactionsApiResult{}

	err = json.Unmarshal(resp.Body(), &apiResult)

	if err != nil {
		return nil, err
	}

	return apiResult.Result, nil
}

