package routes

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/configs/app_config"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/controllers/auth_controller"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/controllers/ping"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	app.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
	// Ping test
	app.GET("/ping", ping.GetPong)

	SetupUserRouter(app)
	SetupBookRouter(app)
	SetupFileRouter(app)

	app.POST("login", auth_controller.Login)
}
