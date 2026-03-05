package app_config

import "os"

var PORT string

func App_Port() {
	PORT = os.Getenv("APP_PORT")
}
