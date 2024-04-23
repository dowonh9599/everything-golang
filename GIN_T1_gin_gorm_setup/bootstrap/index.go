package bootstrap

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs/app_config"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs/cors_config"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/database"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func BootstrapApp() {
	// LOAD .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// INIT server with Gin Framework
	app := gin.Default()

	// INIT CONFIG
	//app.Use(cors_config.CorsConfig)
	app.Use(cors_config.CorsConfigContrib())
	configs.InitConfigs()

	// INIT DB CONNECTION
	database.ConnectDatabase()

	// SETUP endpoint routes
	routes.SetupRouter(app)

	// RUN app at port
	app.Run(app_config.PORT)
}
