package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/generator"
)

func GeneratorCode(c *gin.Context) {}
func GeneratorCodeByDatabase(c *gin.Context) {
	results, err := generator.GeneratorByDatabase()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(results, c)
}
