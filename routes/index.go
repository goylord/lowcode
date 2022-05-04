package routes

import "github.com/gin-gonic/gin"

func RouterRegister(engine *gin.Engine) {
	GeneratorRouterRegistry(engine)
	ControlsFieldMapRouterRegistry(engine)
	ControlsTypeRouterRegistry(engine)
	DictRouterRegistry(engine)
	DictItemRouterRegistry(engine)
	FieldRouterRegistry(engine)
	FieldTypeRouterRegistry(engine)
	TableEntityRouterRegistry(engine)
	UserRouterRegistry(engine)
}
