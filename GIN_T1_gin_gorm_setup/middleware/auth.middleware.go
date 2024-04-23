package middleware

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if !strings.Contains(bearerToken, "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token is invalid",
		})

		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token must be provided",
		})

		return
	}

	claimsData, errDecode := utils.DecodeToken(token)
	if errDecode != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})

		return
	}

	c.Set("claimsData", claimsData)
	c.Set("user_id", claimsData["id"])
	c.Set("email", claimsData["email"])
	c.Set("username", claimsData["username"])
	c.Set("password", claimsData["password"])
	c.Set("created_at", claimsData["created_at"])
	c.Set("updated_at", claimsData["updated_at"])

	c.Next()
}
