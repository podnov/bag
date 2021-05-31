package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/podnov/bag/api"
	"github.com/podnov/bag/api/bscscan"
	"github.com/podnov/bag/api/coinmarketcap"
	"github.com/podnov/bag/api/ox"
)

func CheckLiveness(c *gin.Context) {
	c.Status(http.StatusOK)
}

func CheckReadiness(c *gin.Context) {
	c.Status(http.StatusOK)
}

func CheckRoot(c *gin.Context) {
	c.Status(http.StatusOK)
}

func createDataFetcher(cryptocurrencyMapStore *coinmarketcap.CryptocurrencyMapStore) api.DataFetcher {
	bscClient := &bscscan.BscApiClient{}
	oxClient := &ox.OxApiClient{}

	return api.NewDataFetcher(bscClient,
		cryptocurrencyMapStore,
		oxClient,
	)
}

func GetAccount(c *gin.Context, cryptocurrencyMapStore *coinmarketcap.CryptocurrencyMapStore) {
	accountId := c.Param("accountId")
	dataFetcher := createDataFetcher(cryptocurrencyMapStore)

	statistics, err := dataFetcher.GetAccountStatistics(accountId)

	var responseStatus int
	var responseBody interface{}

	if err == nil {
		responseStatus = http.StatusOK
		responseBody = statistics
	} else {
		responseStatus = http.StatusInternalServerError
		fmt.Printf("Error: %v", fmt.Sprint(err))
		// TODO better
	}

	c.JSON(responseStatus, responseBody)
}
