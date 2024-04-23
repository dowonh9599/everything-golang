package file_controller

import (
	"fmt"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/constant"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"path/filepath"
)

func HandleUploadFile(c *gin.Context) {
	claimsData := c.MustGet("claimsData").(jwt.MapClaims)
	fmt.Printf("user email", claimsData["email"])

	userId := c.MustGet("user_id").(float64)
	fmt.Printf("user id", userId)

	fileHeader, _ := c.FormFile("file")
	if fileHeader == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "file is required",
		})

		return
	}

	allowedFileTypes := []string{"image/png", "image/jpg", "application/pdf"}
	isValid := utils.ValidateFile(fileHeader, allowedFileTypes)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "file type now allowed",
		})

		return
	}

	fileExt := filepath.Ext(fileHeader.Filename)
	filename := utils.GetRandomFilename(fileExt)

	errSave := utils.SaveFile(c, fileHeader, &filename)
	if errSave != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errSave.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "file uploaded",
	})
}

func HandleRemoveFile(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "filename is required",
		})

		return
	}

	errRemove := utils.RemoveFile(constant.DIR_FILE + "/" + filename)
	if errRemove != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errRemove.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "file successfully deleted",
	})
}

func SendStatus(c *gin.Context) {
	filename := c.MustGet("filename").(string)

	c.JSON(http.StatusOK, gin.H{
		"message":   "file uploaded",
		"file_name": filename,
	})
}
