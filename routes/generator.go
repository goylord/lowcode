package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func GeneratorRouterRegistry(engine *gin.Engine) {
	group := engine.Group("generator")
	{
		group.POST("database", services.GeneratorCodeByDatabase)
	}
}
