package routes

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/controllers/book_controller"
	"github.com/gin-gonic/gin"
)

func SetupBookRouter(app *gin.Engine) {
	app.GET("/books", book_controller.GetAllBooks)
}
