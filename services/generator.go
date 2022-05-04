package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/generator"
)

func GeneratorCode(c *gin.Context) {}
func GeneratorCodeByDatabase(c *gin.Context) {
	generator.GeneratorByDatabase()
	common.ResultSuccess("success", c)
}
