package bscscan

import (
	"os"
	"testing"
)

func Test_formatAccountTokenBalanceUrl(t *testing.T) {
	givenAccountAddress := "given-account-address"
	givenApiKey := "given-api-key"
	givenTokenAddress := "given-token-address"

	os.Setenv(ApiKeyEnvironmentVariableName, givenApiKey)

	client := BscApiClient{}
	actual := client.formatAccountTokenBalanceUrl(givenAccountAddress, givenTokenAddress)

	expected := "https://api.bscscan.com/api?module=account&action=tokenbalance&address=given-account-address&contractaddress=given-token-address&tag=latest&apikey=given-api-key"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}

func Test_formatAccountTokenTransactionsUrl(t *testing.T) {
	givenAccountAddress := "given-account-address"
	givenApiKey := "given-api-key"

	os.Setenv(ApiKeyEnvironmentVariableName, givenApiKey)

	client := BscApiClient{}
	actual := client.formatAccountTokenTransactionsUrl(givenAccountAddress)

	expected := "https://api.bscscan.com/api?module=account&action=tokentx&address=given-account-address&sort=asc&apikey=given-api-key"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}

func Test_formatAccountTokenTransactionsForTokenUrl(t *testing.T) {
	givenAccountAddress := "given-account-address"
	givenApiKey := "given-api-key"
	givenTokenAddress := "given-token-address"

	os.Setenv(ApiKeyEnvironmentVariableName, givenApiKey)

	client := BscApiClient{}
	actual := client.formatAccountTokenTransactionsForTokenUrl(givenAccountAddress, givenTokenAddress)

	expected := "https://api.bscscan.com/api?module=account&action=tokentx&address=given-account-address&contractaddress=given-token-address&sort=asc&apikey=given-api-key"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}

