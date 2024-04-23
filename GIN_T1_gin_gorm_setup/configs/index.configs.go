package configs

import (
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs/app_config"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs/db_config"
	"github.com/dowonh9599/everything-series/everything-golang/GIN_T1_gin_gorm_setup/configs/log_config"
)

func InitConfigs() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
	log_config.InitLoggingConfig(app_config.LOG_FILE_PATH)
	// and more...
}
