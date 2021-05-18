package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/podnov/bag/api/controllers"
)

var urlSchemePattern = regexp.MustCompile(`^(?P<Scheme>https?)://.+$`)

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
		var scheme string

		if headerValues, _ := c.Request.Header["Origin"]; len(headerValues) > 0 {
			origin := headerValues[0]
			scheme = extractUrlScheme(origin)
		} else {
			scheme = "https"
		}

		allowOrigin := fmt.Sprintf("%s://cryptobag.podnov.com", scheme)

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func extractUrlScheme(url string) (string) {
	matches := urlSchemePattern.FindStringSubmatch(url)
	return  matches[1]
}
