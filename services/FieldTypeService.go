package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/databases"
	"github.com/goylold/lowcode/models"
)

const FieldTypeTableName = "FieldType"

func FieldTypeList(c *gin.Context) {
	engine := databases.GetXormEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.FieldType
	err = requestParams.DisposeRequest(engine.Table(FieldTypeTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(FieldTypeTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func FieldTypeAdd(c *gin.Context) {
	var table models.FieldType
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

func FieldTypeUpdate(c *gin.Context) {
	var table models.FieldType
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

func FieldTypeDelete(c *gin.Context) {
	var table models.FieldType
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

func FieldTypeGetOne(c *gin.Context) {
	var table models.FieldType
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
