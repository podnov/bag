package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/podnov/bag/api/controllers"
)

func main() {
	r := gin.Default()

	corsConfig := createCorsConfig()
	r.Use(corsConfig)

	r.GET("/", controllers.CheckRoot)
	r.GET("/bag/api/v1/accounts/:accountId", controllers.GetAccount) 
	r.GET("/bag/api/v1/health/liveness", controllers.CheckLiveness)
	r.GET("/bag/api/v1/health/readiness", controllers.CheckReadiness)

	r.Run()
}

func createCorsConfig() (gin.HandlerFunc) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://*.podnov.com",
		"httpx://*.podnov.com",
	}

	return cors.New(config)
}
