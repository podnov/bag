package ox

import (
	"testing"
)

func Test_formatTokenUrl(t *testing.T) {
	client := OxApiClient{}

	givenTokenAddress := "given-token-address"

	actual := client.formatQuoteUrl(givenTokenAddress)

	expected := "https://bsc.api.0x.org/swap/v1/quote?buyToken=BUSD&sellToken=given-token-address&sellAmount=1000000000000000000&excludedSources=BakerySwap,Belt,DODO,DODO_V2,Ellipsis,Mooniswap,MultiHop,Nerve,SushiSwap,Smoothy,ApeSwap,CafeSwap,CheeseSwap,JulSwap,LiquidityProvider&slippagePercentage=0&gasPrice=0&intentOnFilling=false"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}
