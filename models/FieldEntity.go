package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/databases"
	"github.com/goylold/lowcode/utils"
)

const FieldTableName string = "Field"

type Field struct {
  
  Code string `form:"code" json:"code"`
  
  ControlsType string `form:"controlsType" json:"controlsType"`
  
  CreatedBy string `form:"created_by" json:"created_by"`
  
  CreatedTime string `form:"created_time" json:"created_time"`
  
  DefaultValue string `form:"defaultValue" json:"defaultValue"`
  
  DictCode string `form:"dictCode" json:"dictCode"`
  
  Id string `form:"id" json:"id"`
  
  IsNull string `form:"isNull" json:"isNull"`
  
  IsSort string `form:"isSort" json:"isSort"`
  
  IsSync string `form:"isSync" json:"isSync"`
  
  Length string `form:"length" json:"length"`
  
  Name string `form:"name" json:"name"`
  
  Revision int64 `form:"revision" json:"revision"`
  
  TableId string `form:"tableId" json:"tableId"`
  
  Type string `form:"type" json:"type"`
  
  UpdatedBy string `form:"updated_by" json:"updated_by"`
  
  UpdatedTime string `form:"updated_time" json:"updated_time"`
  
}

func (table *Field) Add() error {
	engine := databases.GetXormEngine()
	table.Id = utils.GetUUID()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	table.CreatedBy = common.GetCurrentUser()
	table.UpdatedBy = common.GetCurrentUser()
	_, err := engine.Table(FieldTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Field) Update() error {
	engine := databases.GetXormEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedBy = common.GetCurrentUser()
	_, err := engine.Table(FieldTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Field) Delete() error {
	engine := databases.GetXormEngine()
	affected, err := engine.Table(FieldTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *Field) GetOne(id string) error {
	engine := databases.GetXormEngine()
	_, err := engine.Table(FieldTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}