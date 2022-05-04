package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func FieldTypeRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/FieldType")
	{
		group.POST("/list", services.FieldTypeList)
		group.POST("/add", services.FieldTypeAdd)
		group.POST("/update", services.FieldTypeUpdate)
		group.POST("/delete/:id", services.FieldTypeDelete)
		group.POST("/get/:id", services.FieldTypeGetOne)
	}
}
