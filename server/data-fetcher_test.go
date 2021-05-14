package server

import "math/big"
import "testing"

import "github.com/podnov/bag/server/bscscan"

func Test_calculateEarnedRawTokens(t *testing.T) {
	givenAccountAddress := "given-account-address"
	givenSwapAddress := "given-swap-address"
	givenBalance := big.NewInt(42000)

	givenTransactions := []bscscan.TransactionApiResult {
		bscscan.TransactionApiResult {
			From: givenSwapAddress,
			Value: "101", // buy
			To: givenAccountAddress,
		},
		bscscan.TransactionApiResult {
			From: givenAccountAddress,
			Value: "202", // sell
			To: givenSwapAddress,
		},
		bscscan.TransactionApiResult {
			From: givenSwapAddress,
			Value: "303", // buy
			To: givenAccountAddress,
		},
	}

	actual, err := calculateEarnedRawTokens(givenAccountAddress, givenBalance, givenTransactions)

	if err != nil {
		t.Errorf("Got %s; want nil", err)
	}

	expected := big.NewInt(41798)

	if actual.Cmp(expected) != 0 {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}

func Test_parseBigInt_fail(t *testing.T) {
	givenValue := "abc"

	_, actual := parseBigInt(givenValue)

	if actual == nil {
		t.Errorf("Got nil; want non-nil")
	}

	actualMessage := actual.Error()
	expectedMessage := "Could not parse [abc] as big.Int"

	if actualMessage != expectedMessage {
		t.Errorf("Got %s; want %s", actualMessage, expectedMessage)
	}
}

func Test_parseBigInt_success(t *testing.T) {
	givenValue := "4242"

	actual, err := parseBigInt(givenValue)

	if err != nil {
		t.Errorf("Got %s; want nil", err)
	}

	expected := big.NewInt(4242)

	if actual.Cmp(expected) != 0 {
		t.Errorf("Got %s; want %s", actual, expected)
	}
}
