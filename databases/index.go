package databases

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var driverName string = "mysql"
var dataSourceName string = "root:123456@tcp(127.0.0.1:3307)/lowcode"
var dataSourceInfoName string = "root:123456@tcp(127.0.0.1:3307)/information_schema"

var engine *xorm.Engine
var engine2 *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Panic("loading database source fail!")
	}
	engine.ShowSQL(true)

	engine2, err = xorm.NewEngine(driverName, dataSourceInfoName)
	if err != nil {
		log.Panic("loading database source 2 fail!")
	}
	engine2.ShowSQL(true)
}

func GetXormEngine() *xorm.Engine {
	return engine
}

func GetTableInfoEngine() *xorm.Engine {
	return engine2
}
