package bootstrap

import (
	"simple-audit-log-system/app/config"
	"simple-audit-log-system/app/config/app_config"
	"simple-audit-log-system/app/controller"
	"simple-audit-log-system/app/database"
	"simple-audit-log-system/app/repository"
	"simple-audit-log-system/app/router"
	"simple-audit-log-system/app/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() {
	_ = godotenv.Load()

	config.Init_Config()
	database.Connect()

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"MESSAGE": "Application Is Running Well 🌸",
		})
	})

	repository := repository.NewCriminalRepositoryRegistry()
	service := service.NewCriminalServiceRegistry(repository)
	controller := controller.NewCriminalControllerRegistry(service)

	router.CriminalRouter(app, controller)

	app.Run(app_config.PORT)
}
