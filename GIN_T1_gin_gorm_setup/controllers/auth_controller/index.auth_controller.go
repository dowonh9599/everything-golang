package auth_controller

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/database"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/models"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/requests"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func Login(c *gin.Context) {
	loginReq := new(requests.LoginRequest)
	if errReq := c.ShouldBind(&loginReq); errReq != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	user := new(models.User)
	// find user by email passed through loginReq
	errUser := database.DB.Table("users").Where("email = ?", loginReq.Email).Find(&user).Error
	if errUser != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credentials are invalid",
		})

		return
	}

	// check password
	if loginReq.Password != *user.Password {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credentials are invalid",
		})

		return
	}

	// create claims
	claims := jwt.MapClaims{
		"id":         user.Id,
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["username"] = user.Username
	claims["password"] = user.Password
	claims["created_at"] = user.CreatedAt
	claims["updated_at"] = user.UpdatedAt

	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generated token",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}
