package main

import (
	"generator/create"
	"generator/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 初始化数据库
	db.InitDB(db.DbConfig{
		UserName:  "root",
		PassWord:  "root",
		DbHost:    "127.0.0.1",
		DbPort:    3306,
		DbName:    "ynd_common",
		DbCharset: "utf8",
	})

	// 需要生成的数据库
	tableNames := []string{"sys_user", "sys_admin"}

	// 创建模版
	create.Generator.InitTemplateConfig(create.InitConfig{
		PackageName: "common",
		TabNames:    tableNames,
		Output:      "build",
	}).CreateProtos("proto.gohtml").TmplSample()
}
