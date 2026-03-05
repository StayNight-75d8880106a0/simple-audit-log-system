package db_config

import "os"

type DB_Config_Init struct {
	DB_HOST     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
	DB_PORT     string
	DB_SSLMODE  string
	DB_TIMEZONE string
}

func DB_Config() *DB_Config_Init {
	return &DB_Config_Init{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
		DB_TIMEZONE: os.Getenv("DB_TIMEZONE"),
	}
}
