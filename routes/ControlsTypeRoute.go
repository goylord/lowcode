package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func ControlsTypeRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/ControlsType")
	{
		group.POST("/list", services.ControlsTypeList)
		group.POST("/add", services.ControlsTypeAdd)
		group.POST("/update", services.ControlsTypeUpdate)
		group.POST("/delete/:id", services.ControlsTypeDelete)
		group.POST("/get/:id", services.ControlsTypeGetOne)
	}
}
