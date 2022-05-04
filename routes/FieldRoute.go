package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func FieldRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/Field")
	{
		group.POST("/list", services.FieldList)
		group.POST("/add", services.FieldAdd)
		group.POST("/update", services.FieldUpdate)
		group.POST("/delete/:id", services.FieldDelete)
		group.POST("/get/:id", services.FieldGetOne)
	}
}
