package utils

import (
	"fmt"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs/app_config"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

var charSet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	str := make([]byte, n)
	for i := range str {
		str[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(str)
}

func GetRandomFilename(ext string, prefixes ...string) string {
	prefix := "file"
	if len(prefixes) != 0 {
		prefix = prefixes[0]
	}

	timeNow := time.Now().UTC().Format("20061206")

	return fmt.Sprintf("%s-%s-%s%s", prefix, timeNow, RandomString(5), ext)
}

func ValidateFile(fileHeader *multipart.FileHeader, allowedFileTypes []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	for _, allowedFileType := range allowedFileTypes {
		if contentType == allowedFileType {
			return true
		}
	}
	return false
}

func SaveFile(c *gin.Context, fileHeader *multipart.FileHeader, filename *string) error {
	errUpload := c.SaveUploadedFile(fileHeader, fmt.Sprintf("%s/files/%s", app_config.STATIC_DIR, *filename))
	if errUpload != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return errUpload
	}

	return nil
}

func RemoveFile(filepath string) error {
	err := os.Remove(filepath)

	if err != nil {
		log.Println("failed to remove file")
		return err
	}
	return nil
}
