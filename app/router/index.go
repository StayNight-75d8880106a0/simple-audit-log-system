package router

import (
	"simple-audit-log-system/app/controller"

	"github.com/gin-gonic/gin"
)

func CriminalRouter(app *gin.Engine, CriminalController *controller.Criminal_Controller) {

	r := app.Group("/api/criminal")

	r.POST("/create", CriminalController.Create)
	r.PUT("/update/:id", CriminalController.Update)
	r.DELETE("/delete/:id", CriminalController.Delete)
	r.GET("/", CriminalController.GetAll)
	r.GET("/audit", CriminalController.GetAllAudit)

}
