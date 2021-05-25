package coinmarketcap

import (
	"encoding/json"
	"testing"
)

const givenCryptocurrencyMapJson = `{
    "status": {
        "timestamp": "2021-05-24T22:58:33.496Z",
        "error_code": 0,
        "error_message": null,
        "elapsed": 18,
        "credit_count": 1,
        "notice": null
    },
    "data": [
        {
            "id": 1,
            "name": "Bitcoin",
            "symbol": "BTC",
            "slug": "bitcoin",
            "rank": 1,
            "is_active": 1,
            "first_historical_data": "2013-04-28T18:47:21.000Z",
            "last_historical_data": "2021-05-24T22:49:03.000Z",
            "platform": null
        },
        {
            "id": 2,
            "name": "Litecoin",
            "symbol": "LTC",
            "slug": "litecoin",
            "rank": 13,
            "is_active": 1,
            "first_historical_data": "2013-04-28T18:47:22.000Z",
            "last_historical_data": "2021-05-24T22:49:02.000Z",
            "platform": null
        },
        {
            "id": 3,
            "name": "Namecoin",
            "symbol": "NMC",
            "slug": "namecoin",
            "rank": 680,
            "is_active": 1,
            "first_historical_data": "2013-04-28T18:47:22.000Z",
            "last_historical_data": "2021-05-24T22:49:02.000Z",
            "platform": null
        },
        {
            "id": 4,
            "name": "Terracoin",
            "symbol": "TRC",
            "slug": "terracoin",
            "rank": 1764,
            "is_active": 1,
            "first_historical_data": "2013-04-28T18:47:22.000Z",
            "last_historical_data": "2021-05-24T22:49:02.000Z",
            "platform": null
        }
    ]
}`

func Test_convertCryptocurrencyMap(t *testing.T) {
	given := CryptocurrencyMapApiResult{}

	err := json.Unmarshal([]byte(givenCryptocurrencyMapJson), &given)

	if err != nil {
		t.Errorf("Got %s; want nil", err)
	}

	actual := convertCryptocurrencyMap(given)

	actualBytes, err := json.MarshalIndent(actual, "", "	")

	if err != nil {
		t.Errorf("Got %s; want nil", err)
	}

	actualJson := string(actualBytes)
	expectedJson := `{
	"BTC": {
		"id": 1,
		"name": "Bitcoin",
		"symbol": "BTC",
		"slug": "bitcoin",
		"rank": 1,
		"is_active": 1,
		"first_historical_data": "2013-04-28T18:47:21Z",
		"last_historical_data": "2021-05-24T22:49:03Z",
		"platform": {
			"id": 0,
			"name": "",
			"symbol": "",
			"slug": "",
			"token_address": ""
		}
	},
	"LTC": {
		"id": 2,
		"name": "Litecoin",
		"symbol": "LTC",
		"slug": "litecoin",
		"rank": 13,
		"is_active": 1,
		"first_historical_data": "2013-04-28T18:47:22Z",
		"last_historical_data": "2021-05-24T22:49:02Z",
		"platform": {
			"id": 0,
			"name": "",
			"symbol": "",
			"slug": "",
			"token_address": ""
		}
	},
	"NMC": {
		"id": 3,
		"name": "Namecoin",
		"symbol": "NMC",
		"slug": "namecoin",
		"rank": 680,
		"is_active": 1,
		"first_historical_data": "2013-04-28T18:47:22Z",
		"last_historical_data": "2021-05-24T22:49:02Z",
		"platform": {
			"id": 0,
			"name": "",
			"symbol": "",
			"slug": "",
			"token_address": ""
		}
	},
	"TRC": {
		"id": 4,
		"name": "Terracoin",
		"symbol": "TRC",
		"slug": "terracoin",
		"rank": 1764,
		"is_active": 1,
		"first_historical_data": "2013-04-28T18:47:22Z",
		"last_historical_data": "2021-05-24T22:49:02Z",
		"platform": {
			"id": 0,
			"name": "",
			"symbol": "",
			"slug": "",
			"token_address": ""
		}
	}
}`

	if actualJson != expectedJson {
		t.Errorf("Got %s; want %s", actualJson, expectedJson)
	}
}

