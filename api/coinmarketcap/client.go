package coinmarketcap

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/podnov/bag/api/utils"
)

const ApiKeyEnvironmentVariableName = "COIN_MARKET_CAP_API_KEY"
const apiKeyHeaderName = "X-CMC_PRO_API_KEY"

const apiBaseUrl = "https://pro-api.coinmarketcap.com/v1"

// TODO handle client invocation failures

type CmcApiClient struct {

}

func (c *CmcApiClient) createRestyClient() (*resty.Client) {
	return resty.New()
}

func (c *CmcApiClient) formatCryptocurrencyMapUrl() (string) {
	return fmt.Sprintf("%s/cryptocurrency/map", apiBaseUrl)
}

func (c *CmcApiClient) GetCryptocurrencyMap() (CryptocurrencyMapApiResult, error) {
	apiKey := utils.GetEnv(ApiKeyEnvironmentVariableName)
	url := c.formatCryptocurrencyMapUrl()

	client := c.createRestyClient()

	resp, err := client.R().
		SetHeader(apiKeyHeaderName, apiKey).
		Get(url)

	if err != nil {
		return CryptocurrencyMapApiResult{}, err
	}

	result := CryptocurrencyMapApiResult{}

	err = json.Unmarshal(resp.Body(), &result)

	if err != nil {
		return CryptocurrencyMapApiResult{}, err
	}

	return result, nil
}

