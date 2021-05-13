package pancakeswap

import "testing"

func Test_formatTokenUrl(t *testing.T) {
	client := PancakeswapApiClient{}

	givenTokenAddress := "given-token-address"

	actual := client.formatTokenUrl(givenTokenAddress)

	expected := "https://api.pancakeswap.info/api/v2/tokens/given-token-address"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}

