package server

import "math/big"
import "testing"

import "github.com/podnov/bag/server/bscscan"

func Test_calculateEarnedRawTokens(t *testing.T) {
	givenBalance := big.NewInt(42000)

	givenTransactions := []bscscan.TransactionApiResult {
		bscscan.TransactionApiResult {
			Value: "101",
		},
		bscscan.TransactionApiResult {
			Value: "202",
		},
		bscscan.TransactionApiResult {
			Value: "303",
		},
	}

	actual, err := calculateEarnedRawTokens(givenBalance, givenTransactions)

	if err != nil {
		t.Errorf("Got %s; want nil", err)
	}

	expected := big.NewInt(41394)

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
