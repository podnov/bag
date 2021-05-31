package ox

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const apiBaseUrl = "https://bsc.api.0x.org/swap/v1"
const excludedSources = "BakerySwap,Belt,DODO,DODO_V2,Ellipsis,Mooniswap,MultiHop,Nerve,SushiSwap,Smoothy,ApeSwap,CafeSwap,CheeseSwap,JulSwap,LiquidityProvider"

// TODO handle client invocation failures

type OxApiClient struct {
}

func (c *OxApiClient) createRestyClient() *resty.Client {
	return resty.New()
}

func (c *OxApiClient) formatQuoteUrl(tokenAddress string) string {
	return fmt.Sprintf("%s/quote?buyToken=BUSD&sellToken=%s&sellAmount=1000000000000000000&excludedSources=%s&slippagePercentage=0&gasPrice=0&intentOnFilling=false", apiBaseUrl, tokenAddress, excludedSources)
}

func (c *OxApiClient) GetQuote(tokenAddress string) (*QuoteApiResult, error) {
	url := c.formatQuoteUrl(tokenAddress)

	client := c.createRestyClient()

	resp, err := client.R().
		Get(url)

	if err != nil {
		return nil, err
	}

	apiResult := QuoteApiResult{}

	err = json.Unmarshal(resp.Body(), &apiResult)

	if err != nil {
		return nil, err
	}

	return &apiResult, nil
}
