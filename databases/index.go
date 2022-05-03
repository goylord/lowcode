package databases

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var driverName string = "mysql"
var dataSourceName string = "root:123456@tcp(127.0.0.1:3307)/lowcode"

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Panic("loading database source fail!")
	}
	engine.ShowSQL(true)
}

func GetXormEngine() *xorm.Engine {
	return engine
}
