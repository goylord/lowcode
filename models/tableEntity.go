package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/databases"
	"github.com/goylold/lowcode/utils"
)

const tableName string = "TableEntity"

type TableEntity struct {
	CreatedBy   string `form:"created_by" json:"created_by"`
	CreatedTime string `form:"created_time" json:"created_time"`
	UpdatedBy   string `form:"updated_by" json:"updated_by"`
	UpdatedTime string `form:"updated_time" json:"updated_time"`
	Id          string `form:"id" json:"id"`
	Name        string `form:"name" json:"name" binding:"required"`
	Code        string `form:"code" json:"code" binding:"required"`
	Remark      string `form:"remark" json:"remark"`
}

func (table *TableEntity) Add() error {
	engine := databases.GetXormEngine()
	table.Id = utils.GetUUID()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	table.CreatedBy = common.GetCurrentUser()
	table.UpdatedBy = common.GetCurrentUser()
	_, err := engine.Table(tableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *TableEntity) Update() error {
	engine := databases.GetXormEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedBy = common.GetCurrentUser()
	_, err := engine.Table(tableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *TableEntity) Delete() error {
	engine := databases.GetXormEngine()
	affected, err := engine.Table(tableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}
