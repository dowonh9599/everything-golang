package book_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "books",
	})
}
