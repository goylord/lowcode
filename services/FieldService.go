package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/databases"
	"github.com/goylold/lowcode/models"
)

const FieldTableName = "Field"

func FieldList(c *gin.Context) {
	engine := databases.GetXormEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.Field
	err = requestParams.DisposeRequest(engine.Table(FieldTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(FieldTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func FieldAdd(c *gin.Context) {
	var table models.Field
	err := c.ShouldBindJSON(&table)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	err = table.Add()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}

func FieldUpdate(c *gin.Context) {
	var table models.Field
	err := c.ShouldBindJSON(&table)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	if table.Id == "" {
		common.ResultError(500, "Id不能为空", c)
		return
	}
	err = table.Update()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}

func FieldDelete(c *gin.Context) {
	var table models.Field
	id := c.Param("id")
	if id == "" {
		common.ResultError(500, "Id不能为空", c)
		return
	}
	table.Id = id
	err := table.Delete()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}

func FieldGetOne(c *gin.Context) {
	var table models.Field
	id := c.Param("id")
	if id == "" {
		common.ResultError(500, "Id不能为空", c)
		return
	}
	err := table.GetOne(id)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}
