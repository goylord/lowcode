package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func TableEntityRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/tableEntity")
	{
		group.POST("/list", services.TableEntityList)
		group.POST("/add", services.TableEntityLAdd)
		group.POST("/update", services.TableEntityUpdate)
		group.POST("/delete/:id", services.TableEntityDelete)
	}
}
