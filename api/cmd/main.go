package main

import (
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
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://*.podnov.com,https://*podnov.com")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
