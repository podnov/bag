package controllers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/podnov/bag/api"
	"github.com/podnov/bag/api/bscscan"
	pcsv1 "github.com/podnov/bag/api/pancakeswap/v1"
	pcsv2 "github.com/podnov/bag/api/pancakeswap/v2"
)
var oneHundred = big.NewFloat(float64(100))

func CheckLiveness(c *gin.Context) {
	c.Status(http.StatusOK)
}

func CheckReadiness(c *gin.Context) {
	c.Status(http.StatusOK)
}

func CheckRoot(c *gin.Context) {
	c.Status(http.StatusOK)
}

func createDataFetcher() (api.DataFetcher) {
	bscClient := &bscscan.BscApiClient{}
	pcsv1Client := &pcsv1.PancakeswapApiClient{}
	pcsv2Client := &pcsv2.PancakeswapApiClient{}

	return api.NewDataFetcher(bscClient, pcsv1Client, pcsv2Client)
}

func GetAccount(c *gin.Context) {
	accountId := c.Param("accountId")
	dataFetcher := createDataFetcher()

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

