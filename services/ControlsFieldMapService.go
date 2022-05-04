package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/databases"
	"github.com/goylold/lowcode/models"
)

const ControlsFieldMapTableName = "ControlsFieldMap"

func ControlsFieldMapList(c *gin.Context) {
	engine := databases.GetXormEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.ControlsFieldMap
	err = requestParams.DisposeRequest(engine.Table(ControlsFieldMapTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(ControlsFieldMapTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func ControlsFieldMapAdd(c *gin.Context) {
	var table models.ControlsFieldMap
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

func ControlsFieldMapUpdate(c *gin.Context) {
	var table models.ControlsFieldMap
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

func ControlsFieldMapDelete(c *gin.Context) {
	var table models.ControlsFieldMap
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

func ControlsFieldMapGetOne(c *gin.Context) {
	var table models.ControlsFieldMap
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
