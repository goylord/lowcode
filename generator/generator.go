package generator

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/goylold/lowcode/databases"
	"github.com/goylold/lowcode/utils"
)

// 代码生成器

type TableInfo struct {
	TableName string
	Fields    []Field
}

type Field struct {
	StructName string
	FieldType  string
	SourceName string
	Required   bool
	Conditions string
}

var TypeMap map[string]string = map[string]string{
	"varchar":  "string",
	"int":      "int64",
	"datetime": "string",
}

func GeneratorCode() {

}

func GeneratorByDatabase() (string, error) {
	engine := databases.GetXormEngine()
	results, err := engine.SQL("show tables").QueryString()
	var routes []string
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	engineTableInfo := databases.GetTableInfoEngine()
	for _, tableInfoName := range results {
		tableName := tableInfoName["Tables_in_lowcode"]
		if tableName == "" {
			continue
		}
		modelFilePath := "./models/" + tableName + "Entity.go"
		serviceFilePath := "./services/" + tableName + "Service.go"
		routeFilePath := "./routes/" + tableName + "Route.go"
		if utils.IsExists(modelFilePath) || utils.IsExists(serviceFilePath) || utils.IsExists(routeFilePath) {
			continue
		}
		fields, err := engineTableInfo.SQL("select * from columns where table_name = ?", tableName).QueryString()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		TableInfos := &TableInfo{
			TableName: utils.Ucfirst(tableName),
		}
		for _, field := range fields {
			fieldName := field["COLUMN_NAME"]
			fieldType, ok := TypeMap[field["DATA_TYPE"]]
			if !ok {
				fieldType = "string"
			}
			isRequired := false
			conditions := ""
			if field["COLUMN_NAME"] == "NO" && fieldName != "id" {
				isRequired = true
				conditions = ` binding:"required"`
			}
			fieldStruct := Field{
				StructName: utils.Ucfirst(utils.Case2Camel(fieldName)),
				SourceName: fieldName,
				FieldType:  fieldType,
				Required:   isRequired,
				Conditions: conditions,
			}
			TableInfos.Fields = append(TableInfos.Fields, fieldStruct)
		}

		// log.Println(TableInfos)
		modelFileTmpl, err := ioutil.ReadFile("./template/model.tmpl")
		if err != nil {
			log.Println("model模板文件打开失败", err.Error())
		}
		serviceFileTmpl, err := ioutil.ReadFile("./template/service.tmpl")
		if err != nil {
			log.Println("service模板文件打开失败", err.Error())
		}
		routeFileTmpl, err := ioutil.ReadFile("./template/route.tmpl")
		if err != nil {
			log.Println("route模板文件打开失败", err.Error())
		}
		modelOutputFile, err := os.OpenFile(modelFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println("无法成功打开文件", modelFilePath, err.Error())
		}
		serviceOutputFile, err := os.OpenFile(serviceFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println("无法成功打开文件", serviceFilePath, err.Error())
		}
		routeOutputFile, err := os.OpenFile(routeFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println("无法成功打开文件", routeFilePath, err.Error())
		}

		defer modelOutputFile.Close()
		defer serviceOutputFile.Close()
		defer routeOutputFile.Close()

		if err != nil {
			log.Println(err.Error())
			return "", err
		}

		tpl, err := template.New("model").Parse(string(modelFileTmpl))
		if err != nil {
			log.Println(err.Error())
		}
		err = tpl.Execute(modelOutputFile, TableInfos)
		if err != nil {
			log.Println(err.Error())
		}

		tplService, err := template.New("service").Parse(string(serviceFileTmpl))
		if err != nil {
			log.Println(err.Error())
		}
		err = tplService.Execute(serviceOutputFile, TableInfos)
		if err != nil {
			log.Println(err.Error())
		}

		tplRoute, err := template.New("route").Parse(string(routeFileTmpl))
		if err != nil {
			log.Println(err.Error())
		}
		err = tplRoute.Execute(routeOutputFile, TableInfos)

		if err != nil {
			log.Println(err.Error())
		}
		routes = append(routes, TableInfos.TableName+"RouterRegistry(engine)")
	}
	return strings.Join(routes, ";"), nil
}
