package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/databases"
	"github.com/goylold/lowcode/models"
)

const DictTableName = "Dict"

func DictList(c *gin.Context) {
	engine := databases.GetXormEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.Dict
	err = requestParams.DisposeRequest(engine.Table(DictTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(DictTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func DictAdd(c *gin.Context) {
	var table models.Dict
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

func DictUpdate(c *gin.Context) {
	var table models.Dict
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

func DictDelete(c *gin.Context) {
	var table models.Dict
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

func DictGetOne(c *gin.Context) {
	var table models.Dict
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
