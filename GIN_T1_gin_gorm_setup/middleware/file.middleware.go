package middleware

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func UploadFile(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")

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

	c.Set("filename", filename)

	c.Next()
}
