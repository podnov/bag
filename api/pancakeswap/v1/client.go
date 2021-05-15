package v1

import "encoding/json"
import "fmt"

import "github.com/go-resty/resty/v2"

const apiBaseUrl = "https://api.pancakeswap.info/api"

// TODO handle client invocation failures

type PancakeswapApiClient struct {

}

func (c *PancakeswapApiClient) createRestyClient() (*resty.Client) {
	return resty.New()
}

func (c *PancakeswapApiClient) formatTokenUrl(tokenAddress string) (string) {
	return fmt.Sprintf("%s/tokens/%s", apiBaseUrl, tokenAddress)
}

func (c *PancakeswapApiClient) GetToken(tokenAddress string) (*TokenApiResult, error) {
	url := c.formatTokenUrl(tokenAddress)

	client := c.createRestyClient()

	resp, err := client.R().
		Get(url)

	if err != nil {
		return nil, err
	}

	apiResult := TokenApiResult{}

	err = json.Unmarshal(resp.Body(), &apiResult)

	if err != nil {
		return nil, err
	}

	return &apiResult, nil
}

