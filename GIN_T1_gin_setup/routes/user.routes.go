package routes

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_setup/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(app *gin.Engine) {
	userRoute := app.Group("user")

	app.GET("/users", user_controller.GetAllUsers)
	userRoute.GET("/:id", user_controller.GetUserById)
	userRoute.GET("/paginate", user_controller.GetUsersPaginate) // /user/paginate?perPage=2*page=2
	userRoute.POST("/", user_controller.AddNewUser)
	userRoute.PATCH("/:id", user_controller.UpdateById)
	userRoute.DELETE("/:id", user_controller.DeleteById)
}
