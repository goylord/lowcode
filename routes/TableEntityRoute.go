package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func TableEntityRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/TableEntity")
	{
		group.POST("/list", services.TableEntityList)
		group.POST("/add", services.TableEntityAdd)
		group.POST("/update", services.TableEntityUpdate)
		group.POST("/delete/:id", services.TableEntityDelete)
		group.POST("/get/:id", services.TableEntityGetOne)
	}
}
