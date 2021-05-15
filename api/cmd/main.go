package main

import (
	"github.com/gin-gonic/gin"
	"github.com/podnov/bag/api/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/bag/api/v1/accounts/:accountId", controllers.GetAccount) 
	r.GET("/bag/api/v1/health/liveness", controllers.CheckLiveness)
	r.GET("/bag/api/v1/health/readiness", controllers.CheckReadiness)

	r.Run()
}

