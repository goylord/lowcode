package routes

import "github.com/gin-gonic/gin"

func RouterRegister(engine *gin.Engine) {
	TableEntityRouterRegistry(engine)
}
