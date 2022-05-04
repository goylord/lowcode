package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func DictRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/Dict")
	{
		group.POST("/list", services.DictList)
		group.POST("/add", services.DictAdd)
		group.POST("/update", services.DictUpdate)
		group.POST("/delete/:id", services.DictDelete)
		group.POST("/get/:id", services.DictGetOne)
	}
}
