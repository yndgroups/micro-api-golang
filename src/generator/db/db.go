package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type DbConfig struct {
	UserName  string `json:"userName"`
	PassWord  string `json:"passWord"`
	DbHost    string `json:"dbHost"`
	DbName    string `json:"dbName"`
	DbPort    uint16 `json:"dbProt"`
	DbCharset string `json:"charset"`
}

func InitDB(dbConfig DbConfig) {
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.UserName,
		dbConfig.PassWord,
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbName,
		dbConfig.DbCharset)
	if db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		panic("数据连接错误 info => " + err.Error())
	} else {
		DB = db
	}
}

// 查询表字段
func QueryTableField[T any](tableName string) []T {
	var list = make([]T, 0)
	DB.Raw(`SHOW FULL COLUMNS FROM ` + tableName).Scan(&list)
	return list
}
