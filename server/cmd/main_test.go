package main

import "testing"

func Test_formatAccountTokenBalanceUrl(t *testing.T) {
	var givenAccountAddress = "given-account-address"
	var givenTokenAddress = "given-token-address"

	var actual = formatAccountTokenBalanceUrl(givenAccountAddress, givenTokenAddress)

	var expected = "https://api.bscscan.com/api?module=account&action=tokenbalance&address=given-account-address&contractaddress=given-token-address&tag=latest&apikey=X5T7BC9KUVXAWP6SQHRD5Z7RXRH58RJVIX"

	if actual != expected {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}
