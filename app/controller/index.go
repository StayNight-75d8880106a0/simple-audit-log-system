package controller

import (
	"simple-audit-log-system/app/helper"
	"simple-audit-log-system/app/interface/service_interface"
	"simple-audit-log-system/app/request"

	"github.com/gin-gonic/gin"
)

type Criminal_Controller struct {
	service service_interface.Service_Interface
}

func NewCriminalControllerRegistry(criminal_service service_interface.Service_Interface) *Criminal_Controller {
	return &Criminal_Controller{
		service: criminal_service,
	}
}

func (c *Criminal_Controller) Create(ctx *gin.Context) {

	request := new(request.Criminal_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	criminal, errCreate := c.service.Create(request)

	if errCreate != nil {
		if appError, ok := errCreate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errCreate.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errCreate.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"Message": "Succes",
		"Data":    criminal,
	})

}

func (c *Criminal_Controller) GetAll(ctx *gin.Context) {

	criminal, errGet := c.service.GetAll()

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errGet.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success",
		"Data":    criminal,
	})

}

func (c *Criminal_Controller) Update(ctx *gin.Context) {

	ID := ctx.Param("id")

	request := new(request.Criminal_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	criminal, errCreate := c.service.Update(ID, request)

	if errCreate != nil {
		if appError, ok := errCreate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errCreate.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errCreate.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"Message": "Succes",
		"Data":    criminal,
	})

}

func (c *Criminal_Controller) Delete(ctx *gin.Context) {

	ID := ctx.Param("id")

	errDelete := c.service.Delete(ID)

	if errDelete != nil {
		if appError, ok := errDelete.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errDelete.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errDelete.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success",
	})

}

func (c *Criminal_Controller) GetAllAudit(ctx *gin.Context) {

	audit, errGet := c.service.GetAllAudit()

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errGet.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success",
		"Data":    audit,
	})

}
