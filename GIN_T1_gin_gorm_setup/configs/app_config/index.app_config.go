package app_config

import "os"

var PORT = ":8080"
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"
var SECRET_KEY = "SECRET_KEY"
var LOG_FILE_PATH = "logs/file/gin.log"

func InitAppConfig() {
	port := os.Getenv("PORT")
	if port != "" {
		PORT = port
	}

	staticRoute := os.Getenv("STATIC_ROUTE")
	if staticRoute != "" {
		STATIC_ROUTE = staticRoute
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir != "" {
		STATIC_DIR = staticDir
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey != "" {
		SECRET_KEY = secretKey
	}

	logFilePath := os.Getenv("LOG_FILE_PATH")
	if logFilePath != "" {
		LOG_FILE_PATH = logFilePath
	}
}
