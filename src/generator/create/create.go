package create

import (
	"fmt"
	"generator/db"
	"generator/model"
	"html/template"
	"os"
	"strings"
)

var PackageName string

type InitConfig struct {
	PackageName string   `json:"packageName"`
	TabNames    []string `json:"tabName"`
	Output string ``
}

var Generator = &generator{}

type generator struct {
	templates   *template.Template
	tables      []string
	packageName string
}

// 初始化模版
func (g *generator) InitTemplateConfig(config InitConfig) *generator {
	// 初始化模版
	g.templates = template.Must(template.ParseGlob("templates/*.gohtml"))
	g.templates.Funcs(createFuncMap())
	g.tables = config.TabNames
	g.packageName = config.PackageName

	return g
}

// 处理mysql 字段和go的对应关系
func (g *generator) getFieldType(typeName string) string {
	fType := "string"
	mySqlType := make(map[string][]string)
	mySqlType["string"] = []string{"varchar", "char", "tinytext", "text", "mediumtext", "longtext", "date", "time", "datetime", "timestamp", "year"}
	mySqlType["int64"] = []string{"tinyint", "smallint", "mediumint", "int", "bigint"}
	mySqlType["float64"] = []string{"float", "double", "decimal"}
	// mySqlType["date"] = []string{"date", "time", "datetime", "timestamp", "year"}
	mySqlType["enum"] = []string{"enum"}
	mySqlType["set"] = []string{"set"}
	mySqlType["blob"] = []string{"binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob"}

L:
	for k, v := range mySqlType {
		if len(v) > 0 {
			for _, c := range v {
				if strings.Contains(typeName, strings.ToLower(c)) {
					fType = k
					break L
				}
			}
		}
	}
	return fType
}

// 创建proto文件
func (g *generator) CreateProtos(templateNames string) *generator {
	for _, s := range g.tables {
		g.CreateProto(s)
	}
	return g
}

func (g *generator) CreateProto(tableName string) {
	list := db.QueryTableField[model.TableField](tableName)
	if len(list) <= 0 {
		return
	}
	var fList = make([]model.TableField, 0)
	for _, v := range list {
		v.Type = g.getFieldType(v.Type)
		v.Field = UderscoreToUpperCamelCase(v.Field)
		fList = append(fList, v)
		println(fmt.Sprintf("%v", v))
	}
	file, err := os.Create("build/" + tableName + ".proto")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// template, err := Template.ParseFiles("templates/" + templateNames)
	data := make(map[string]any, 0)
	data["modelName"] = UderscoreToUpperCamelCase(tableName)
	data["list"] = fList
	data["packageName"] = g.packageName
	g.templates.ExecuteTemplate(file, "proto.gohtml", data)
}

func (g *generator) TmplSample() *generator {
	file, err := os.Create("build/example.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	g.templates.ExecuteTemplate(file, "example.gohtml", nil)
	return g
}
