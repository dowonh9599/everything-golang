package log_config

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
)

func createLogsFolderIfNotExist(path string) {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("Creating", dir, "directory...")
		err := os.MkdirAll(dir, 0755)

		if err != nil {
			log.Println("Fail to create", dir)
		} else {
			log.Println("creating", dir, "directory done.")
		}
	}
}

func openOrCreateLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		var errCreateFile error
		logFile, errCreateFile = os.Create(path)

		if errCreateFile != nil {
			log.Println("Can't create log file", errCreateFile)

		}
	}

	return logFile, nil
}

func InitLoggingConfig(path string) {
	gin.DisableConsoleColor()

	createLogsFolderIfNotExist(path)
	f, _ := openOrCreateLogFile(path)

	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)

}
