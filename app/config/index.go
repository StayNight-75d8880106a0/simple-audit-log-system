package config

import (
	"simple-audit-log-system/app/config/app_config"
	"simple-audit-log-system/app/config/db_config"
)

func Init_Config() {
	app_config.App_Port()
	db_config.DB_Config()
}
