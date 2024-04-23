package user_controller

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/database"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/models"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/requests"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user := new(responses.UserReponse)
	errDb := database.DB.Table("users").Where("id = ?", id).First(&user).Error

	if errDb != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": "internal server error",
		})
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUsersPaginate(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}

	perPage := c.Query("perPage")
	if perPage == "" {
		perPage = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)

	users := new([]models.User)

	errDb := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error

	if errDb != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     users,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}

func GetAllUsers(c *gin.Context) {
	users := new([]responses.UserReponse)
	// query users data from database
	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func AddNewUser(c *gin.Context) {
	userReq := new(requests.UserRequest)

	// validate binding
	if errReq := c.ShouldBind(&userReq); errReq != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	user := new(models.User)
	user.Id = &userReq.Id
	user.Username = &userReq.Username
	user.Email = &userReq.Email
	user.Password = &userReq.Password

	errDb := database.DB.Table("users").Create(&user).Error

	if errDb != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errDb.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data saved successfully",
		"data":    user,
	})
}

func UpdateById(c *gin.Context) {
	id := c.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)

	if errReq := c.ShouldBind(&userReq); errReq != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	// find user from database
	if errDb := database.DB.Table("users").Where("id = ?", id).First(&user).Error; errDb != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errDb.Error(),
		})

		return
	}
	if user.Id == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})

		return
	}

	// update data
	timeNow := time.Now()
	user.Id = &userReq.Id
	user.Username = &userReq.Username
	user.Email = &userReq.Email
	user.Password = &userReq.Password
	user.UpdatedAt = &timeNow

	errUpdate := database.DB.Table("users").Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errUpdate.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user data updated",
		"data":    user,
	})
}

func DeleteById(c *gin.Context) {
	id := c.Param("id")

	// find if user exists
	userResp := new(responses.UserReponse)
	errFindUser := database.DB.Table("users").Where("id = ?", id).First(userResp).Error
	if errFindUser != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})

		return
	}

	errDb := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&models.User{}).Error
	if errDb != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user data deleted successfully",
	})
}
