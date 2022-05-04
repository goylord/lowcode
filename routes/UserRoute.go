package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func UserRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/User")
	{
		group.POST("/list", services.UserList)
		group.POST("/add", services.UserAdd)
		group.POST("/update", services.UserUpdate)
		group.POST("/delete/:id", services.UserDelete)
		group.POST("/get/:id", services.UserGetOne)
	}
}
