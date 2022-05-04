package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func ControlsFieldMapRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/ControlsFieldMap")
	{
		group.POST("/list", services.ControlsFieldMapList)
		group.POST("/add", services.ControlsFieldMapAdd)
		group.POST("/update", services.ControlsFieldMapUpdate)
		group.POST("/delete/:id", services.ControlsFieldMapDelete)
		group.POST("/get/:id", services.ControlsFieldMapGetOne)
	}
}
