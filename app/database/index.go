package database

import (
	"fmt"
	"simple-audit-log-system/app/config/db_config"
	"simple-audit-log-system/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB_Cluster struct {
	DB *gorm.DB
}

var dbCluster *DB_Cluster

func Connect() *DB_Cluster {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().DB_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE)

	db, ErrorConnect := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if ErrorConnect != nil {
		panic("An error occurred while trying to connect to the database!! : " + ErrorConnect.Error())
	}

	errMigrate := db.AutoMigrate(&model.CriminalFugitives{}, &model.AuditLogs{})

	if errMigrate != nil {
		panic("Failed to Migrate Database!! : " + errMigrate.Error())
	}

	fmt.Println("=========================================")
	fmt.Println("🚀 Database Cluster Status:")
	fmt.Println("✅ Database Connection: OK!")
	fmt.Println("✅ Auto Migration: Success!")
	fmt.Println("=========================================")

	dbCluster = &DB_Cluster{
		DB: db,
	}

	return dbCluster
}

func GetInstanceDbCluster() *DB_Cluster {
	return dbCluster
}
