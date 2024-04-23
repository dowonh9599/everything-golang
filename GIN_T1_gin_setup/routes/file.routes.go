package routes

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/controllers/file_controller"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/middleware"
	"github.com/gin-gonic/gin"
)

func SetupFileRouter(app *gin.Engine) {
	authRoute := app.Group("file", middleware.AuthMiddleware)

	// file controllers
	authRoute.POST("/", file_controller.HandleUploadFile)
	authRoute.POST("/middleware", middleware.UploadFile, file_controller.SendStatus)
	authRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
}
