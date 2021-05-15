package bscscan

import "testing"

func Test_formatAccountTokenBalanceUrl(t *testing.T) {
	client := BscApiClient{}

	givenAccountAddress := "given-account-address"
	givenTokenAddress := "given-token-address"

	actual := client.formatAccountTokenBalanceUrl(givenAccountAddress, givenTokenAddress)

	expected := "https://api.bscscan.com/api?module=account&action=tokenbalance&address=given-account-address&contractaddress=given-token-address&tag=latest&apikey=X5T7BC9KUVXAWP6SQHRD5Z7RXRH58RJVIX"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}

func Test_formatAccountTokenTransactionsUrl(t *testing.T) {
	client := BscApiClient{}

	givenAccountAddress := "given-account-address"

	actual := client.formatAccountTokenTransactionsUrl(givenAccountAddress)

	expected := "https://api.bscscan.com/api?module=account&action=tokentx&address=given-account-address&sort=asc&apikey=X5T7BC9KUVXAWP6SQHRD5Z7RXRH58RJVIX"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}

func Test_formatAccountTokenTransactionsForTokenUrl(t *testing.T) {
	client := BscApiClient{}

	givenAccountAddress := "given-account-address"
	givenTokenAddress := "given-token-address"

	actual := client.formatAccountTokenTransactionsForTokenUrl(givenAccountAddress, givenTokenAddress)

	expected := "https://api.bscscan.com/api?module=account&action=tokentx&address=given-account-address&contractaddress=given-token-address&sort=asc&apikey=X5T7BC9KUVXAWP6SQHRD5Z7RXRH58RJVIX"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}
