package cors_config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsConfig(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTION, PUT, PATCH, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}

func CorsConfigContrib() gin.HandlerFunc {
	var allowedOrigins = []string{
		"*",
	}
	var allowedHeaders = []string{
		"Content-Type, Content-Length",
	}
	var allowedMethods = []string{"GET, POST, OPTION, PUT, PATCH, DELETE"}

	config := cors.DefaultConfig()

	config.AllowOrigins = allowedOrigins
	config.AllowCredentials = true
	config.AllowHeaders = allowedHeaders
	config.AllowMethods = allowedMethods

	return cors.New(config)
}
