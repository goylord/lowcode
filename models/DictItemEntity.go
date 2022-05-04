package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/databases"
	"github.com/goylold/lowcode/utils"
)

const DictItemTableName string = "DictItem"

type DictItem struct {
  
  CreatedBy string `form:"created_by" json:"created_by"`
  
  CreatedTime string `form:"created_time" json:"created_time"`
  
  Id string `form:"id" json:"id"`
  
  Name string `form:"name" json:"name"`
  
  Remark string `form:"remark" json:"remark"`
  
  Revision int64 `form:"revision" json:"revision"`
  
  UpdatedBy string `form:"updated_by" json:"updated_by"`
  
  UpdatedTime string `form:"updated_time" json:"updated_time"`
  
  Value string `form:"value" json:"value"`
  
}

func (table *DictItem) Add() error {
	engine := databases.GetXormEngine()
	table.Id = utils.GetUUID()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	table.CreatedBy = common.GetCurrentUser()
	table.UpdatedBy = common.GetCurrentUser()
	_, err := engine.Table(DictItemTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *DictItem) Update() error {
	engine := databases.GetXormEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedBy = common.GetCurrentUser()
	_, err := engine.Table(DictItemTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *DictItem) Delete() error {
	engine := databases.GetXormEngine()
	affected, err := engine.Table(DictItemTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *DictItem) GetOne(id string) error {
	engine := databases.GetXormEngine()
	_, err := engine.Table(DictItemTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}