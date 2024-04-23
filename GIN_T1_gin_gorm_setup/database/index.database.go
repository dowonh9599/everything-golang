package database

import (
	"fmt"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs/db_config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func ConnectDatabase() {
	var errConnection error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db_config.DB_HOST, db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_NAME, db_config.DB_PORT)
	db, errConnection := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errConnection != nil {
		panic("Connection to Postgres failed.")
	} else {
		DB = db
		fmt.Println("Successfully connected database")
	}
}
