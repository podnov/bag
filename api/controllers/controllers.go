package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/podnov/bag/api"
	"github.com/podnov/bag/api/bscscan"
	"github.com/podnov/bag/api/coinmarketcap"
	pcsv1 "github.com/podnov/bag/api/pancakeswap/v1"
	pcsv2 "github.com/podnov/bag/api/pancakeswap/v2"
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
	pcsv1Client := &pcsv1.PancakeswapApiClient{}
	pcsv2Client := &pcsv2.PancakeswapApiClient{}

	return api.NewDataFetcher(bscClient,
		cryptocurrencyMapStore,
		pcsv1Client,
		pcsv2Client,
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
