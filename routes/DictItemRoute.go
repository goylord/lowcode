package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func DictItemRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/DictItem")
	{
		group.POST("/list", services.DictItemList)
		group.POST("/add", services.DictItemAdd)
		group.POST("/update", services.DictItemUpdate)
		group.POST("/delete/:id", services.DictItemDelete)
		group.POST("/get/:id", services.DictItemGetOne)
	}
}
